package sqlite

import (
	"context"

	"github.com/brittonhayes/aos/api"
	"github.com/uptrace/bun"
)

func (r *repository) citiesFilterQuery(query *bun.SelectQuery, params api.GetCitiesParams) (*bun.SelectQuery, error) {

	if params.Name != nil {
		query = query.Where("? LIKE ?", bun.Ident("name"), *params.Name+"%")
	}

	return query, nil
}

func (r *repository) GetCities(ctx context.Context, params api.GetCitiesParams) ([]api.City, error) {
	var cities []api.City

	query, err := r.citiesFilterQuery(r.db.NewSelect().OrderExpr("id ASC").Model(&cities), params)
	if err != nil {
		return nil, err
	}

	err = query.Scan(ctx, &cities)
	if err != nil {
		return nil, err
	}

	return cities, nil
}

func (r *repository) GetCityByID(ctx context.Context, id string) (*api.City, error) {
	var city api.City

	err := r.db.NewSelect().Model(&city).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &city, nil
}
