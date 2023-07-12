package sqlite

import (
	"context"

	"github.com/brittonhayes/warhammer/api"
)

func (r *warhammerRepository) GetAllegianceByID(ctx context.Context, id string) (*api.Allegiance, error) {
	var allegiance api.Allegiance
	err := r.db.NewSelect().Model(&allegiance).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &allegiance, nil
}

func (r *warhammerRepository) GetAllegiances(ctx context.Context) ([]api.Allegiance, error) {
	var allegiances []api.Allegiance
	err := r.db.NewSelect().Model(&allegiances).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return allegiances, nil
}
