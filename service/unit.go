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
	var unit api.Unit

	if unit.ID == nil {
		return api.GetUnitByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrUnitNotFound})
	}

	return api.GetUnitByIDJSON200Response(unit)
}

func (s *WarhammerService) GetUnits(w http.ResponseWriter, r *http.Request) *api.Response {
	var units []api.Unit

	units = append(units, api.Unit{
		Name: "Unit 1",
	})

	if len(units) == 0 {
		return api.GetUnitsJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrUnitsNotFound})
	}

	return api.GetUnitsJSON200Response(units)
}

func (s *WarhammerService) CreateUnit(w http.ResponseWriter, r *http.Request) *api.Response {
	var unit api.Unit
	return api.CreateUnitJSON201Response(unit)
}

func (s *WarhammerService) UpdateUnitByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {
	var unit api.Unit
	return api.UpdateUnitByIDJSON200Response(unit)
}

func (s *WarhammerService) DeleteUnitByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {
	var unit api.Unit
	return api.DeleteUnitByIDJSON200Response(unit)
}
