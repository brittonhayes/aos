package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/brittonhayes/aos"
	"github.com/brittonhayes/aos/fixtures"
	"github.com/brittonhayes/aos/internal/logging"
	"github.com/brittonhayes/aos/server"
	"github.com/brittonhayes/aos/service"
	"github.com/brittonhayes/aos/sqlite"

	"github.com/urfave/cli/v2"
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
			config := &server.Config{
				Repository:  sqlite.NewRepository(c.String("db")),
				Service:     service.New(repo),
				ServiceName: c.String("service"),
				Environment: c.String("environment"),
				Port:        c.Int("port"),
				Tracing:     c.Bool("tracing"),
			}

			s := server.New(c.Context, config)

			logging.NewLogrus(c.Context).Infof("Starting server on port %d", config.Port)
			return s.ListenAndServe()
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
