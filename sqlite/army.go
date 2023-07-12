package sqlite

import (
	"context"

	"github.com/brittonhayes/warhammer/api"
)

func (r *warhammerRepository) GetArmyByID(ctx context.Context, id string) (*api.Army, error) {
	var army api.Army
	err := r.db.NewSelect().Model(&army).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &army, nil
}

func (r *warhammerRepository) GetArmies(ctx context.Context) ([]api.Army, error) {
	var armies []api.Army
	err := r.db.NewSelect().Model(&armies).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return armies, nil
}

func (r *warhammerRepository) UpdateArmyByID(ctx context.Context, id string, army *api.Army) (*api.Army, error) {
	_, err := r.db.NewUpdate().Model(army).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return army, nil
}
