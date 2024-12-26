package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/brittonhayes/aos"
	"github.com/brittonhayes/aos/api"
	"github.com/brittonhayes/aos/graph"
	"github.com/brittonhayes/aos/internal/logging"
	"github.com/brittonhayes/aos/internal/sunset"
	"github.com/brittonhayes/aos/internal/tracing"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	server *http.Server
	Router *chi.Mux
}

type Config struct {
	Repository  aos.Repository
	Service     api.ServerInterface
	ServiceName string
	Environment string
	Port        int
	Tracing     bool
}

func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func New(ctx context.Context, config *Config) *Server {

	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	r := chi.NewRouter()

	// Add sunset middleware before all other middleware
	// to display a sunset notice if the current date is after the sunset date
	r.Use(sunset.Middleware)

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.CleanPath)
	r.Use(middleware.SetHeader("Cache-Control", "public, max-age=86400, s-maxage=31536000, stale-while-revalidate=60"))
	r.Use(middleware.Timeout(4 * time.Second))
	r.Use(logging.Middleware)
	// r.Use(apimw.OAPIValidator(swagger))

	// Use the OpenTelemetry middleware to trace all requests
	if config.Tracing {
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
	api.Handler(config.Service, api.WithRouter(r))
	r.Post("/query", graphQLHandler(config.Repository))
	r.Get("/graphql", apolloHandler("/query"))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(aos.HTML_HOME)
	})

	if config.Tracing {
		traceProvider, err := tracing.InitTracer(config.ServiceName, config.Environment)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := traceProvider.Shutdown(ctx); err != nil {
				log.Printf("Error shutting down tracer provider: %v", err)
			}
		}()

		meterProvider, err := tracing.InitMeter()
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := meterProvider.Shutdown(ctx); err != nil {
				log.Printf("Error shutting down meter provider: %v", err)
			}
		}()
	}

	port := config.Port
	if port == 0 {
		port = 8080
	}

	httpsrv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
	}

	return &Server{
		server: httpsrv,
		Router: r,
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
