// Package warhammer provides the WarhammerRepository interface for the Warhammer REST API and services
//
//go:generate goapi-gen -generate types,server,spec -package api --out api/warhammer.gen.go ./api/openapi.yaml
package warhammer

import (
	"context"
	"io/fs"

	"github.com/brittonhayes/warhammer/api"
)

type WarhammerRepository interface {
	GetUnitByID(ctx context.Context, id string) (*api.Unit, error)
	GetUnits(ctx context.Context) ([]api.Unit, error)

	GetArmyByID(ctx context.Context, id string) (*api.Army, error)
	GetArmies(ctx context.Context) ([]api.Army, error)

	GetGrandAllianceByID(ctx context.Context, id string) (*api.GrandAlliance, error)
	GetGrandAlliances(ctx context.Context) ([]api.GrandAlliance, error)

	GetAllegianceByID(ctx context.Context, id string) (*api.Allegiance, error)
	GetAllegiances(ctx context.Context, params api.GetAllegiancesParams) ([]api.Allegiance, error)

	GetGrandStrategyByID(ctx context.Context, id string) (*api.GrandStrategy, error)
	GetGrandStrategies(ctx context.Context) ([]api.GrandStrategy, error)

	GetCities(ctx context.Context, params api.GetCitiesParams) ([]api.City, error)

	Init(ctx context.Context) error
	Generate(ctx context.Context, name string) error
	Migrate(ctx context.Context) error
	Lock(ctx context.Context) error
	Unlock(ctx context.Context) error
	Rollback(ctx context.Context) error
	Seed(ctx context.Context, fsys fs.FS, names ...string) error
}
