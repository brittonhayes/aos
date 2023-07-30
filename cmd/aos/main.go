package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/brittonhayes/aos"
	"github.com/brittonhayes/aos/api"
	"github.com/brittonhayes/aos/fixtures"
	"github.com/brittonhayes/aos/graph"
	"github.com/brittonhayes/aos/internal/logging"
	"github.com/brittonhayes/aos/internal/tracing"
	"github.com/brittonhayes/aos/service"
	"github.com/brittonhayes/aos/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/urfave/cli/v2"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

func main() {

	var (
		repo aos.Repository
	)

	app := &cli.App{
		Name:    "aos",
		Usage:   "AoS api server",
		Suggest: true,
		Before: func(c *cli.Context) error {
			if c.Bool("migrate") {
				repo = sqlite.NewRepository(c.String("db"))

				err := repo.Init(c.Context)
				if err != nil {
					return err
				}

				err = repo.Migrate(c.Context)
				if err != nil {
					return err
				}

				data := []string{
					fixtures.ALLEGIANCES,
					fixtures.ALLIANCES,
					fixtures.CITIES,
					fixtures.STRATEGIES,
					fixtures.UNITS,
					fixtures.WARSCROLLS,
				}

				err = repo.Seed(c.Context, aos.FIXTURES, data...)
				if err != nil {
					return err
				}
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			repo = sqlite.NewRepository(c.String("db"))

			// Create an instance of our handler which satisfies the generated interface
			s := service.New(repo)

			swagger, err := api.GetSwagger()
			if err != nil {
				return err
			}

			// Clear out the servers array in the swagger spec, that skips validating
			// that server names match. We don't know how this thing will be run.
			swagger.Servers = nil

			// Create a new chi router
			r := chi.NewRouter()

			// Use our validation middleware to check all requests against the
			// OpenAPI schema.
			r.Use(middleware.Recoverer)
			r.Use(middleware.RedirectSlashes)
			r.Use(middleware.Timeout(5 * time.Second))
			r.Use(logging.Middleware)
			// r.Use(apimw.OAPIValidator(swagger))

			// Use the OpenTelemetry middleware to trace all requests
			if c.Bool("tracing") {
				r.Use(func(h http.Handler) http.Handler {
					return otelhttp.NewHandler(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							h.ServeHTTP(w, r)

							routePattern := chi.RouteContext(r.Context()).RoutePattern()

							span := trace.SpanFromContext(r.Context())
							span.SetName(routePattern)
							span.SetAttributes(semconv.HTTPTarget(r.URL.String()), semconv.HTTPRoute(routePattern))

							labeler, ok := otelhttp.LabelerFromContext(r.Context())
							if ok {
								labeler.Add(semconv.HTTPRoute(routePattern))
							}
						}),
						"",
					)
				})
			}

			// We now register our server above as the handler for the interface
			api.Handler(s, api.WithRouter(r))
			r.Post("/query", graphQLHandler(repo))
			r.Get("/graphql", apolloHandler("/query"))
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(aos.HOMEPAGE)
			})

			if c.Bool("tracing") {
				traceProvider, err := tracing.InitTracer(c.String("service"), c.String("environment"))
				if err != nil {
					return err
				}
				defer func() {
					if err := traceProvider.Shutdown(c.Context); err != nil {
						log.Printf("Error shutting down tracer provider: %v", err)
					}
				}()

				meterProvider, err := tracing.InitMeter()
				if err != nil {
					return err
				}
				defer func() {
					if err := meterProvider.Shutdown(c.Context); err != nil {
						log.Printf("Error shutting down meter provider: %v", err)
					}
				}()
			}

			port := c.Int("port")
			if port == 0 {
				port = 8080
			}

			httpsrv := &http.Server{
				Handler: r,
				Addr:    fmt.Sprintf("0.0.0.0:%d", port),
			}

			logging.NewLogrus(c.Context).Infof("Starting server on port %d", port)
			return httpsrv.ListenAndServe()
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   8080,
				EnvVars: []string{"PORT"},
				Usage:   "port to listen on",
			},
			&cli.StringFlag{
				Name:    "service",
				Aliases: []string{"s"},
				Value:   "aos",
				EnvVars: []string{"SERVICE_NAME"},
				Usage:   "customize the service name",
			},
			&cli.StringFlag{
				Name:    "environment",
				Value:   "development",
				EnvVars: []string{"ENVIRONMENT"},
				Usage:   "set the environment (development or production)",
			},
			&cli.StringFlag{
				Name:    "db",
				Value:   "file:aos.db",
				EnvVars: []string{"DATABASE_URL"},
				Usage:   "database url",
			},
			&cli.BoolFlag{
				Name:    "tracing",
				Aliases: []string{"t"},
				Value:   false,
				EnvVars: []string{"TRACING"},
				Usage:   "enable tracing",
			},
			&cli.BoolFlag{
				Name:    "migrate",
				Aliases: []string{"m"},
				Value:   false,
				EnvVars: []string{"MIGRATE"},
				Usage:   "run migrations before starting the server",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "migrate the database",
				Action: func(c *cli.Context) error {
					return repo.Migrate(c.Context)
				},
				Before: func(c *cli.Context) error {
					repo = sqlite.NewRepository(c.String("db"))
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Name:  "init",
						Usage: "initialize the migrations table (only needed once)",
						Action: func(c *cli.Context) error {
							return repo.Init(c.Context)
						},
					},
					{
						Name:  "generate",
						Usage: "generate a new migration file with the given name",
						Action: func(c *cli.Context) error {
							if c.Args().Len() == 0 {
								return fmt.Errorf("missing migration name")
							}
							return repo.Generate(c.Context, c.Args().First())
						},
					},
					{
						Name:  "lock",
						Usage: "lock migrations",
						Action: func(c *cli.Context) error {
							return repo.Lock(c.Context)
						},
					},
					{
						Name:  "unlock",
						Usage: "unlock migrations",
						Action: func(c *cli.Context) error {
							return repo.Unlock(c.Context)
						},
					},
					{
						Name:  "rollback",
						Usage: "rollback the last migration group",
						Action: func(c *cli.Context) error {
							return repo.Rollback(c.Context)
						},
					},
					{
						Name:  "seed",
						Usage: "seed data into the database",
						Action: func(c *cli.Context) error {
							data := []string{
								fixtures.ALLEGIANCES,
								fixtures.ALLIANCES,
								fixtures.CITIES,
								fixtures.STRATEGIES,
								fixtures.UNITS,
								fixtures.WARSCROLLS,
							}

							return repo.Seed(c.Context, aos.FIXTURES, data...)
						},
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func graphQLHandler(repository aos.Repository) http.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Repo: repository,
	}}))

	h.Use(extension.FixedComplexityLimit(50))
	h.Use(extension.Introspection{})
	return h.ServeHTTP
}

func apolloHandler(endpoint string) http.HandlerFunc {
	playgroundHandler := playground.ApolloSandboxHandler("GraphQL Playground", endpoint)
	return playgroundHandler
}
