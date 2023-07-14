package migrations

import (
	"embed"

	"github.com/uptrace/bun/migrate"
)

var Migrations = migrate.NewMigrations(migrate.WithMigrationsDirectory("."))

//go:embed *.go
var migrations embed.FS

func init() {
	if err := Migrations.Discover(migrations); err != nil {
		panic(err)
	}
}
