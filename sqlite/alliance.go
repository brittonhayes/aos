package sqlite

import (
	"context"

	"github.com/brittonhayes/warhammer"
)

func (r *warhammerRepository) GetGrandAllianceByID(ctx context.Context, id string) (*warhammer.GrandAlliance, error) {
	var grandAlliance warhammer.GrandAlliance
	err := r.db.NewSelect().Model(&grandAlliance).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &grandAlliance, nil
}

func (r *warhammerRepository) GetGrandAlliances(ctx context.Context) ([]warhammer.GrandAlliance, error) {
	var grandAlliances []warhammer.GrandAlliance
	err := r.db.NewSelect().Model(&grandAlliances).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return grandAlliances, nil
}
