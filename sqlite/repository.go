package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/brittonhayes/warhammer"
	"github.com/brittonhayes/warhammer/sqlite/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/migrate"
)

type warhammerRepository struct {
	m  *migrate.Migrator
	db *bun.DB
}

func NewWarhammerRepository(connection string) warhammer.WarhammerRepository {
	sqldb, err := sql.Open(sqliteshim.ShimName, connection)
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	migrator := migrate.NewMigrator(db, migrations.Migrations)

	return &warhammerRepository{
		m:  migrator,
		db: db,
	}
}

func (r *warhammerRepository) Init(ctx context.Context) error {
	return r.m.Init(ctx)
}

func (r *warhammerRepository) Generate(ctx context.Context, name string) error {

	name = strings.ReplaceAll(name, " ", "_")
	ms, err := r.m.CreateSQLMigrations(ctx, name)
	if err != nil {
		return err
	}

	for _, m := range ms {
		fmt.Printf("created migration %s (%s)\n", m.Name, m.Path)
	}

	return nil
}

func (r *warhammerRepository) Migrate(ctx context.Context) error {
	err := r.m.Lock(ctx)
	if err != nil {
		return err
	}
	defer r.m.Unlock(ctx)

	group, err := r.m.Migrate(ctx)
	if err != nil {
		return err
	}

	if group.IsZero() {
		fmt.Printf("there are no new migrations to run (database is up to date)\n")
		return nil
	}

	fmt.Printf("migrated to %s\n", group)
	return nil
}
