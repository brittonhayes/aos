package sqlite

import (
	"context"
	"database/sql"
	"io/fs"
	"strings"

	"github.com/brittonhayes/warhammer"
	"github.com/brittonhayes/warhammer/api"
	"github.com/brittonhayes/warhammer/internal/logging"
	"github.com/brittonhayes/warhammer/sqlite/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bunotel"
	"github.com/uptrace/bun/migrate"
)

type warhammerRepository struct {
	m  *migrate.Migrator
	db *bun.DB
}

const (
	DB_NAME = "warhammer"
)

func NewWarhammerRepository(connection string) warhammer.WarhammerRepository {
	sqldb, err := sql.Open(sqliteshim.ShimName, connection)
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bunotel.NewQueryHook(
		bunotel.WithDBName(DB_NAME),
		bunotel.WithFormattedQueries(true),
	))

	migrator := migrate.NewMigrator(db, migrations.Migrations)

	return &warhammerRepository{
		m:  migrator,
		db: db,
	}
}

func (r *warhammerRepository) Init(ctx context.Context) error {
	logging.NewLogrus(ctx).Info("initializing database")
	return r.m.Init(ctx)
}

func (r *warhammerRepository) Generate(ctx context.Context, name string) error {

	name = strings.ReplaceAll(name, " ", "_")
	m, err := r.m.CreateGoMigration(ctx, name)
	if err != nil {
		return err
	}

	logging.NewLogrus(ctx).Infof("created %s", m.Name)

	return nil
}

func (r *warhammerRepository) Migrate(ctx context.Context) error {

	logger := logging.NewLogrus(ctx)

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
		logger.Info("there are no new migrations to run (database is up to date)")
		return nil
	}

	logger.Infof("migrated to %s\n", group)
	return nil
}

func (r *warhammerRepository) Seed(ctx context.Context, fsys fs.FS, names ...string) error {
	logging.NewLogrus(ctx).Info("seeding database")

	r.db.RegisterModel(
		(*api.Army)(nil),
		(*api.GrandAlliance)(nil),
		(*api.Unit)(nil),
		(*api.Allegiance)(nil),
		(*api.GrandStrategy)(nil),
	)

	f := dbfixture.New(r.db, dbfixture.WithRecreateTables())
	return f.Load(ctx, fsys, names...)
}

func (r *warhammerRepository) Lock(ctx context.Context) error {
	logging.NewLogrus(ctx).Info("locking database")
	return r.m.Lock(ctx)
}

func (r *warhammerRepository) Unlock(ctx context.Context) error {
	logging.NewLogrus(ctx).Info("unlocking database")
	return r.m.Unlock(ctx)
}

func (r *warhammerRepository) Rollback(ctx context.Context) error {
	logger := logging.NewLogrus(ctx)

	err := r.m.Lock(ctx)
	if err != nil {
		return err
	}
	defer r.m.Unlock(ctx)

	logger.Info("rolling back migration")
	group, err := r.m.Rollback(ctx)
	if err != nil {
		return err
	}

	if group.IsZero() {
		logger.Info("there are no migrations to rollback")
		return nil
	}

	logger.Infof("rolled back to %s\n", group)
	return nil
}
