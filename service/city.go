package service

import (
	"net/http"

	"github.com/brittonhayes/warhammer/api"
)

const (
	ErrCitiesNotFound = "cities not found"
)

func (s *WarhammerService) GetCities(w http.ResponseWriter, r *http.Request, params api.GetCitiesParams) *api.Response {
	cities, err := s.repo.GetCities(r.Context(), params)
	if err != nil {
		return api.GetCitiesJSON404Response(api.Error{
			Code:    http.StatusNotFound,
			Message: ErrCitiesNotFound,
		})
	}

	return api.GetCitiesJSON200Response(cities)
}
