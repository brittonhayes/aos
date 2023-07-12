package sqlite

import (
	"context"

	"github.com/brittonhayes/warhammer"
)

func (r *warhammerRepository) GetGrandStrategyByID(ctx context.Context, id string) (*warhammer.GrandStrategy, error) {
	var strategy warhammer.GrandStrategy
	err := r.db.NewSelect().Model(&strategy).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &strategy, nil
}

func (r *warhammerRepository) GetGrandStrategies(ctx context.Context) ([]warhammer.GrandStrategy, error) {
	var strategies []warhammer.GrandStrategy
	err := r.db.NewSelect().Model(&strategies).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return strategies, nil
}
