package service

import (
	"net/http"

	"github.com/brittonhayes/warhammer/api"
	"github.com/brittonhayes/warhammer/internal/logging"
	"github.com/sirupsen/logrus"
)

const (
	ErrAllegianceNotFound  = "Allegiance not found with this ID"
	ErrAllegiancesNotFound = "No allegiances found"
)

func (s *WarhammerService) GetAllegianceByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event":         "GetAllegianceByID",
		"allegiance_id": id,
	}).Info()

	allegiance, err := s.repo.GetAllegianceByID(r.Context(), id)
	if err != nil {
		return api.GetAllegianceByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrAllegianceNotFound})
	}

	return api.GetAllegianceByIDJSON200Response(*allegiance)
}

func (s *WarhammerService) GetAllegiances(w http.ResponseWriter, r *http.Request, params api.GetAllegiancesParams) *api.Response {

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event": "GetAllegiances",
	}).Info()

	allegiances, err := s.repo.GetAllegiances(r.Context(), params)
	if err != nil {
		return api.GetAllegiancesJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrAllegiancesNotFound})
	}

	if len(allegiances) == 0 {
		return api.GetAllegiancesJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrAllegiancesNotFound})
	}

	return api.GetAllegiancesJSON200Response(allegiances)
}
