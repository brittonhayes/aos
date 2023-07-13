package service

import (
	"net/http"

	"github.com/brittonhayes/warhammer/api"
)

const (
	ErrUnitNotFound  = "Unit not found with this ID"
	ErrUnitsNotFound = "No units found"
)

func (s *WarhammerService) GetUnitByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {
	unit, err := s.repo.GetUnitByID(r.Context(), id)
	if err != nil {
		return api.GetUnitByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrUnitNotFound})
	}

	return api.GetUnitByIDJSON200Response(*unit)
}

func (s *WarhammerService) GetUnits(w http.ResponseWriter, r *http.Request) *api.Response {
	units, err := s.repo.GetUnits(r.Context())
	if err != nil {
		return api.GetUnitsJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrUnitsNotFound})
	}

	return api.GetUnitsJSON200Response(units)
}
