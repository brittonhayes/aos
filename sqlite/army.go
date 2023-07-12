package sqlite

import (
	"context"

	"github.com/brittonhayes/warhammer"
)

func (r *warhammerRepository) GetArmyByID(ctx context.Context, id string) (*warhammer.Army, error) {
	var army warhammer.Army
	err := r.db.NewSelect().Model(&army).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &army, nil
}

func (r *warhammerRepository) GetArmies(ctx context.Context) ([]warhammer.Army, error) {
	var armies []warhammer.Army
	err := r.db.NewSelect().Model(&armies).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return armies, nil
}
