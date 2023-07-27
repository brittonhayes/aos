package sqlite

import (
	"context"

	"github.com/brittonhayes/warhammer/api"
	"github.com/uptrace/bun"
)

func (r *warhammerRepository) GetUnitByID(ctx context.Context, id string) (*api.Unit, error) {
	var unit api.Unit
	err := r.db.NewSelect().Model(&unit).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &unit, nil
}

func (r *warhammerRepository) unitsFilterQuery(query *bun.SelectQuery, params api.GetUnitsParams) (*bun.SelectQuery, error) {
	if params.Name != nil {
		query = query.Where("? LIKE ?", bun.Ident("name"), *params.Name+"%")
	}

	if params.Points != nil {
		query = query.Where("? LIKE ?", bun.Ident("points"), *params.Points)
	}

	if params.GrandAlliance != nil {
		query = query.Where("? LIKE ?", bun.Ident("grand_alliance"), *params.GrandAlliance+"%")
	}

	if params.GrandStrategy != nil {
		query = query.Where("? LIKE ?", bun.Ident("grand_strategy"), *params.GrandStrategy+"%")
	}

	return query, nil
}

func (r *warhammerRepository) GetUnits(ctx context.Context, params api.GetUnitsParams) ([]api.Unit, error) {
	var units []api.Unit

	query, err := r.unitsFilterQuery(r.db.NewSelect().Model(&units), params)
	if err != nil {
		return nil, err
	}

	err = query.Scan(ctx, &units)
	if err != nil {
		return nil, err
	}

	return units, nil
}

func (r *warhammerRepository) GetAbilitiesForUnitByID(ctx context.Context, id string) ([]api.Ability, error) {
	var unit api.Unit
	err := r.db.NewSelect().Model(&unit).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return unit.Abilities, nil
}
