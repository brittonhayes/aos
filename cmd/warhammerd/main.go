package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/brittonhayes/warhammer/api"
	"github.com/brittonhayes/warhammer/service"
	"github.com/brittonhayes/warhammer/sqlite"
	"github.com/discord-gophers/goapi-gen/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/urfave/cli/v2"
)

func main() {

	// Create an instance of our handler which satisfies the generated interface
	repo := sqlite.NewWarhammerRepository("file:warhammer.db")
	s := service.NewWarhammerService(repo)

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	r := chi.NewRouter()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OAPIValidator(swagger))

	// We now register our server above as the handler for the interface
	api.Handler(s, api.WithRouter(r))

	app := &cli.App{
		Name:    "warhammerd",
		Usage:   "the warhammer api server",
		Suggest: true,
		Action: func(c *cli.Context) error {
			port := c.Int("port")
			if port == 0 {
				port = 8080
			}

			httpsrv := &http.Server{
				Handler: r,
				Addr:    fmt.Sprintf("0.0.0.0:%d", port),
			}
			fmt.Printf("listening on %s\n", httpsrv.Addr)
			return httpsrv.ListenAndServe()
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   8080,
				Usage:   "port to listen on",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "init",
				Action: func(c *cli.Context) error {
					return repo.Init(c.Context)
				},
			},
			{
				Name: "generate",
				Action: func(c *cli.Context) error {
					if c.Args().Len() == 0 {
						return fmt.Errorf("missing migration name")
					}

					return repo.Generate(c.Context, c.Args().First())
				},
			},
			{
				Name: "migrate",
				Action: func(c *cli.Context) error {
					return repo.Migrate(c.Context)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
