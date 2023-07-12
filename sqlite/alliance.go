package sqlite

import (
	"context"

	"github.com/brittonhayes/warhammer/api"
)

func (r *warhammerRepository) GetGrandAllianceByID(ctx context.Context, id string) (*api.GrandAlliance, error) {
	var grandAlliance api.GrandAlliance
	err := r.db.NewSelect().Model(&grandAlliance).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &grandAlliance, nil
}

func (r *warhammerRepository) GetGrandAlliances(ctx context.Context) ([]api.GrandAlliance, error) {
	var grandAlliances []api.GrandAlliance
	err := r.db.NewSelect().Model(&grandAlliances).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return grandAlliances, nil
}
