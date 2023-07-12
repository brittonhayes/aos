package sqlite

import (
	"context"

	"github.com/brittonhayes/warhammer"
)

func (r *warhammerRepository) GetAllegianceByID(ctx context.Context, id string) (*warhammer.Allegiance, error) {
	var allegiance warhammer.Allegiance
	err := r.db.NewSelect().Model(&allegiance).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &allegiance, nil
}

func (r *warhammerRepository) GetAllegiances(ctx context.Context) ([]warhammer.Allegiance, error) {
	var allegiances []warhammer.Allegiance
	err := r.db.NewSelect().Model(&allegiances).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return allegiances, nil
}
