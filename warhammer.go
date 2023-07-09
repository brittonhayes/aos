package warhammer

import (
	"context"

	"github.com/brittonhayes/warhammer/api"
)

type Army api.Army
type Unit api.Unit

type WarhammerRepository interface {
	GetUnitByID(ctx context.Context, id int64) (*Unit, error)
	GetUnits(ctx context.Context) ([]Unit, error)
	GetArmyByID(ctx context.Context, id int64) (*api.Army, error)
	GetArmies(ctx context.Context) ([]api.Army, error)

	Init(ctx context.Context) error
	Generate(ctx context.Context, name string) error
	Migrate(ctx context.Context) error
}
