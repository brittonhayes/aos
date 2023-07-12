package sqlite

import (
	"context"

	"github.com/brittonhayes/warhammer/api"
	"github.com/uptrace/bun"
)

func (r *warhammerRepository) GetAllegianceByID(ctx context.Context, id string) (*api.Allegiance, error) {
	var allegiance api.Allegiance
	err := r.db.NewSelect().Model(&allegiance).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &allegiance, nil
}

func (r *warhammerRepository) allegiancesFilterQuery(query *bun.SelectQuery, params api.GetAllegiancesParams) (*bun.SelectQuery, error) {
	if params.GrandAlliance != nil {
		query = query.Where("grand_alliance = ?", params.GrandAlliance)
	}

	if params.Name != nil {
		query = query.Where("name = ?", params.Name)
	}

	return query, nil
}

func (r *warhammerRepository) GetAllegiances(ctx context.Context, params api.GetAllegiancesParams) ([]api.Allegiance, error) {
	var allegiances []api.Allegiance

	query, err := r.allegiancesFilterQuery(r.db.NewSelect().Model(&allegiances), params)
	if err != nil {
		return nil, err
	}

	err = query.Scan(ctx, &allegiances)
	if err != nil {
		return nil, err
	}

	return allegiances, nil
}
