package sqlite

import (
	"context"

	"github.com/brittonhayes/warhammer/api"
)

func (r *warhammerRepository) GetUnitByID(ctx context.Context, id string) (*api.Unit, error) {
	var unit api.Unit
	err := r.db.NewSelect().Model(&unit).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &unit, nil
}

func (r *warhammerRepository) GetUnits(ctx context.Context) ([]api.Unit, error) {
	var units []api.Unit
	err := r.db.NewSelect().Model(&units).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return units, nil
}
