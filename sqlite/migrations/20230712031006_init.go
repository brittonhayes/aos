package migrations

import (
	"context"
	"fmt"

	"github.com/brittonhayes/aos/api"
	"github.com/uptrace/bun"
)

var models = []interface{}{
	(*api.Army)(nil),
	(*api.Allegiance)(nil),
	(*api.City)(nil),
	(*api.GrandAlliance)(nil),
	(*api.GrandStrategy)(nil),
	(*api.Unit)(nil),
	(*api.Warscroll)(nil),
}

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")

		for _, model := range models {
			_, err := db.NewCreateTable().Model(model).Exec(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")

		for _, model := range models {
			_, err := db.NewDropTable().Model(model).Exec(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
