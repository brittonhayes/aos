package migrations

import (
	"context"
	"fmt"

	"github.com/brittonhayes/warhammer/api"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")

		_, err := db.NewCreateTable().Model((*api.Army)(nil)).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewCreateTable().Model((*api.Unit)(nil)).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewCreateTable().Model((*api.GrandAlliance)(nil)).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewCreateTable().Model((*api.Allegiance)(nil)).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewCreateTable().Model((*api.GrandStrategy)(nil)).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewCreateTable().Model((*api.City)(nil)).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")

		_, err := db.NewDropTable().Model((*api.Army)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewDropTable().Model((*api.Unit)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewDropTable().Model((*api.GrandAlliance)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewDropTable().Model((*api.Allegiance)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewDropTable().Model((*api.GrandStrategy)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewDropTable().Model((*api.City)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
