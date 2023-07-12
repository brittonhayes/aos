// Package warhammer provides the WarhammerRepository interface for the Warhammer REST API and services
//
//go:generate goapi-gen -generate types,server,spec -package api --out api/warhammer.gen.go ./api/openapi.yaml
package warhammer

import (
	"context"
	"io/fs"

	"github.com/brittonhayes/warhammer/api"
)

type (
	Army          api.Army
	Allegiance    api.Allegiance
	GrandAlliance api.GrandAlliance
	GrandStrategy api.GrandStrategy
	Unit          api.Unit
)

type WarhammerRepository interface {
	GetUnitByID(ctx context.Context, id string) (*Unit, error)
	GetUnits(ctx context.Context) ([]Unit, error)

	GetArmyByID(ctx context.Context, id string) (*Army, error)
	GetArmies(ctx context.Context) ([]Army, error)

	GetGrandAllianceByID(ctx context.Context, id string) (*GrandAlliance, error)
	GetGrandAlliances(ctx context.Context) ([]GrandAlliance, error)

	GetAllegianceByID(ctx context.Context, id string) (*Allegiance, error)
	GetAllegiances(ctx context.Context) ([]Allegiance, error)

	GetGrandStrategyByID(ctx context.Context, id string) (*GrandStrategy, error)
	GetGrandStrategies(ctx context.Context) ([]GrandStrategy, error)

	Init(ctx context.Context) error
	Generate(ctx context.Context, name string) error
	Migrate(ctx context.Context) error
	Lock(ctx context.Context) error
	Unlock(ctx context.Context) error
	Rollback(ctx context.Context) error
	LoadFixtures(ctx context.Context, fsys fs.FS, names ...string) error
}
