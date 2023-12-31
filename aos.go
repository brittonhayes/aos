// Package aos provides the Repository interface for the AoS REST API and services
//
//go:generate mockgen -destination=internal/mocks/mock_repository.go -package=mocks . Repository
//go:generate goapi-gen -generate types,server,spec -package api --out api/api.gen.go ./api/openapi.yaml
//go:generate go run github.com/99designs/gqlgen generate
//go:generate npx @redocly/cli build-docs api/openapi.yaml -o web/src/docs.html
package aos

import (
	"context"
	"io/fs"

	"github.com/brittonhayes/aos/api"
)

type Repository interface {
	GetArmyByID(ctx context.Context, id string) (*api.Army, error)
	GetArmies(ctx context.Context) ([]api.Army, error)

	GetAllegianceByID(ctx context.Context, id string) (*api.Allegiance, error)
	GetAllegiances(ctx context.Context, params *api.GetAllegiancesParams) ([]api.Allegiance, error)

	GetCities(ctx context.Context, params *api.GetCitiesParams) ([]api.City, error)
	GetCityByID(ctx context.Context, id string) (*api.City, error)

	GetGrandAllianceByID(ctx context.Context, id string) (*api.GrandAlliance, error)
	GetGrandAlliances(ctx context.Context) ([]api.GrandAlliance, error)

	GetGrandStrategyByID(ctx context.Context, id string) (*api.GrandStrategy, error)
	GetGrandStrategies(ctx context.Context) ([]api.GrandStrategy, error)

	GetUnitByID(ctx context.Context, id string) (*api.Unit, error)
	GetUnits(ctx context.Context, params *api.GetUnitsParams) ([]api.Unit, error)
	GetAbilitiesForUnitByID(ctx context.Context, id string) ([]api.Ability, error)
	GetWeaponsForUnitByID(ctx context.Context, id string) (*api.WeaponsGroup, error)

	GetWarscrollByID(ctx context.Context, id string) (*api.Warscroll, error)
	GetWarscrolls(ctx context.Context, params *api.GetWarscrollsParams) ([]api.Warscroll, error)

	Init(ctx context.Context) error
	Generate(ctx context.Context, name string) error
	Migrate(ctx context.Context) error
	Lock(ctx context.Context) error
	Unlock(ctx context.Context) error
	Rollback(ctx context.Context) error
	Seed(ctx context.Context, fsys fs.FS, names ...string) error
}
