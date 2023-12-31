package sqlite

import (
	"context"

	"github.com/brittonhayes/aos/api"
	"github.com/uptrace/bun"
)

func (r *repository) GetUnitByID(ctx context.Context, id string) (*api.Unit, error) {
	var unit api.Unit
	err := r.db.NewSelect().Model(&unit).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &unit, nil
}

func (r *repository) unitsFilterQuery(query *bun.SelectQuery, params *api.GetUnitsParams) (*bun.SelectQuery, error) {
	if params == nil {
		return query, nil
	}

	if params.Limit != nil {
		query = query.Limit(*params.Limit)
	}

	if params.Offset != nil {
		query = query.Offset(*params.Offset)
	}

	if params.Name != nil {
		query = query.Where("? LIKE ?", bun.Ident("name"), *params.Name+"%")
	}

	if params.Points != nil {
		query = query.Where("? LIKE ?", bun.Ident("points"), *params.Points)
	}

	if params.Allegiance != nil {
		query = query.Where("? LIKE ?", bun.Ident("allegiance"), *params.Allegiance+"%")
	}

	if params.GrandAlliance != nil {
		query = query.Where("? LIKE ?", bun.Ident("grand_alliance"), *params.GrandAlliance+"%")
	}

	if params.GrandStrategy != nil {
		query = query.Where("? LIKE ?", bun.Ident("grand_strategy"), *params.GrandStrategy+"%")
	}

	return query, nil
}

func (r *repository) GetUnits(ctx context.Context, params *api.GetUnitsParams) ([]api.Unit, error) {
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

func (r *repository) GetAbilitiesForUnitByID(ctx context.Context, id string) ([]api.Ability, error) {
	var unit api.Unit
	err := r.db.NewSelect().OrderExpr("id ASC").Model(&unit).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return unit.Abilities, nil
}

func (r *repository) GetWeaponsForUnitByID(ctx context.Context, id string) (*api.WeaponsGroup, error) {
	var unit api.Unit
	err := r.db.NewSelect().Model(&unit).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	group := &api.WeaponsGroup{
		MeleeWeapons:   unit.MeleeWeapons,
		MissileWeapons: unit.MissileWeapons,
	}

	return group, nil
}
