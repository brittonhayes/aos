package service

import (
	"encoding/json"
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

func (s *WarhammerService) CreateUnit(w http.ResponseWriter, r *http.Request) *api.Response {
	var unit api.Unit
	err := json.NewDecoder(r.Body).Decode(&unit)
	if err != nil {
		return api.CreateUnitJSON400Response(api.Error{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return api.CreateUnitJSON201Response(unit)
}

func (s *WarhammerService) UpdateUnitByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {
	var unit api.Unit
	err := json.NewDecoder(r.Body).Decode(&unit)
	if err != nil {
		return api.UpdateUnitByIDJSON404Response(api.Error{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return api.UpdateUnitByIDJSON200Response(unit)
}

func (s *WarhammerService) DeleteUnitByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {
	var unit api.Unit
	err := json.NewDecoder(r.Body).Decode(&unit)
	if err != nil {
		return api.DeleteUnitByIDJSON404Response(api.Error{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return api.DeleteUnitByIDJSON200Response(unit)
}
