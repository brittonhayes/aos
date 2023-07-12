package migrations

import (
	"context"
	"fmt"

	"github.com/brittonhayes/warhammer"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")

		_, err := db.NewCreateTable().Model((*warhammer.Army)(nil)).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewCreateTable().Model((*warhammer.Unit)(nil)).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewCreateTable().Model((*warhammer.GrandAlliance)(nil)).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewCreateTable().Model((*warhammer.Allegiance)(nil)).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewCreateTable().Model((*warhammer.GrandStrategy)(nil)).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")

		_, err := db.NewDropTable().Model((*warhammer.Army)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewDropTable().Model((*warhammer.Unit)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewDropTable().Model((*warhammer.GrandAlliance)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewDropTable().Model((*warhammer.Allegiance)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewDropTable().Model((*warhammer.GrandStrategy)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
