package sqlite

import (
	"context"
	"database/sql"
	"io/fs"
	"strings"

	"github.com/brittonhayes/aos"
	"github.com/brittonhayes/aos/api"
	"github.com/brittonhayes/aos/internal/logging"
	"github.com/brittonhayes/aos/sqlite/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bunotel"
	"github.com/uptrace/bun/migrate"
)

type aosRepository struct {
	m  *migrate.Migrator
	db *bun.DB
}

const (
	DB_NAME        = "aos"
	DEFAULT_LIMIT  = 100
	DEFAULT_OFFSET = 0
)

func NewRepository(connection string) aos.Repository {
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

	return &aosRepository{
		m:  migrator,
		db: db,
	}
}

func (r *aosRepository) Init(ctx context.Context) error {
	logging.NewLogrus(ctx).Info("initializing database")
	return r.m.Init(ctx)
}

func (r *aosRepository) Generate(ctx context.Context, name string) error {

	name = strings.ReplaceAll(name, " ", "_")
	m, err := r.m.CreateGoMigration(ctx, name)
	if err != nil {
		return err
	}

	logging.NewLogrus(ctx).Infof("created %s", m.Name)

	return nil
}

func (r *aosRepository) Migrate(ctx context.Context) error {

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

func (r *aosRepository) Seed(ctx context.Context, fsys fs.FS, names ...string) error {
	logger := logging.NewLogrus(ctx)

	models := []interface{}{
		(*api.Army)(nil),
		(*api.Allegiance)(nil),
		(*api.City)(nil),
		(*api.GrandAlliance)(nil),
		(*api.GrandStrategy)(nil),
		(*api.Unit)(nil),
		(*api.Warscroll)(nil),
	}

	for _, m := range models {
		logger.Infof("registering model %T", m)
	}

	r.db.RegisterModel(models...)

	logger.Info("seeding database")
	f := dbfixture.New(r.db, dbfixture.WithRecreateTables())

	return f.Load(ctx, fsys, names...)
}

func (r *aosRepository) Lock(ctx context.Context) error {
	logging.NewLogrus(ctx).Info("locking database")
	return r.m.Lock(ctx)
}

func (r *aosRepository) Unlock(ctx context.Context) error {
	logging.NewLogrus(ctx).Info("unlocking database")
	return r.m.Unlock(ctx)
}

func (r *aosRepository) Rollback(ctx context.Context) error {
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
