package service

import (
	"net/http"

	"github.com/brittonhayes/aos/api"
)

const (
	ErrUnitNotFound  = "Unit not found with this ID"
	ErrUnitsNotFound = "No units found"
)

func (s *Service) GetUnitByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {
	unit, err := s.repo.GetUnitByID(r.Context(), id)
	if err != nil {
		return api.GetUnitByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrUnitNotFound})
	}

	return api.GetUnitByIDJSON200Response(*unit)
}

func (s *Service) GetUnits(w http.ResponseWriter, r *http.Request, params api.GetUnitsParams) *api.Response {
	units, err := s.repo.GetUnits(r.Context(), &params)
	if err != nil {
		return api.GetUnitsJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrUnitsNotFound})
	}

	return api.GetUnitsJSON200Response(units)
}

func (s *Service) GetAbilitiesForUnitByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {
	abilities, err := s.repo.GetAbilitiesForUnitByID(r.Context(), id)
	if err != nil {
		return api.GetAbilitiesForUnitByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrUnitNotFound})
	}

	return api.GetAbilitiesForUnitByIDJSON200Response(abilities)
}

func (s *Service) GetWeaponsForUnitByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {
	weapons, err := s.repo.GetWeaponsForUnitByID(r.Context(), id)
	if err != nil {
		return api.GetWeaponsForUnitByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrUnitNotFound})
	}

	return api.GetWeaponsForUnitByIDJSON200Response(*weapons)
}
