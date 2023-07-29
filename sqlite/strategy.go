package sqlite

import (
	"context"

	"github.com/brittonhayes/aos/api"
)

func (r *repository) GetGrandStrategyByID(ctx context.Context, id string) (*api.GrandStrategy, error) {
	var strategy api.GrandStrategy
	err := r.db.NewSelect().Model(&strategy).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &strategy, nil
}

func (r *repository) GetGrandStrategies(ctx context.Context) ([]api.GrandStrategy, error) {
	var strategies []api.GrandStrategy
	err := r.db.NewSelect().OrderExpr("id ASC").Model(&strategies).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return strategies, nil
}
