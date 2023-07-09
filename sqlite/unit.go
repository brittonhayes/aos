package sqlite

import (
	"context"

	"github.com/brittonhayes/warhammer"
)

func (r *warhammerRepository) GetUnitByID(ctx context.Context, id int64) (*warhammer.Unit, error) {
	var unit warhammer.Unit
	err := r.db.NewSelect().Model(&unit).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &unit, nil
}

func (r *warhammerRepository) GetUnits(ctx context.Context) ([]warhammer.Unit, error) {
	var units []warhammer.Unit
	err := r.db.NewSelect().Model(&units).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return units, nil
}
