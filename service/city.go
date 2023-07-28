package service

import (
	"net/http"

	"github.com/brittonhayes/aos/api"
)

const (
	ErrCitiesNotFound = "cities not found"
)

func (s *Service) GetCities(w http.ResponseWriter, r *http.Request, params api.GetCitiesParams) *api.Response {
	cities, err := s.repo.GetCities(r.Context(), params)
	if err != nil {
		return api.GetCitiesJSON404Response(api.Error{
			Code:    http.StatusNotFound,
			Message: ErrCitiesNotFound,
		})
	}

	return api.GetCitiesJSON200Response(cities)
}

func (s *Service) GetCityByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {
	city, err := s.repo.GetCityByID(r.Context(), id)
	if err != nil {
		return api.GetCityByIDJSON404Response(api.Error{
			Code:    http.StatusNotFound,
			Message: ErrCitiesNotFound,
		})
	}

	return api.GetCityByIDJSON200Response(*city)
}
