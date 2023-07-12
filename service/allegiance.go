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
	var allegiance api.Allegiance
	allegiance.ID = &id

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event":         "GetAllegianceByID",
		"allegiance_id": id,
	}).Info()

	if allegiance.ID == nil {
		return api.GetAllegianceByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrAllegianceNotFound})
	}

	return api.GetAllegianceByIDJSON200Response(allegiance)
}

func (s *WarhammerService) GetAllegiances(w http.ResponseWriter, r *http.Request) *api.Response {
	var allegiances []api.Allegiance

	allegiances = append(allegiances, api.Allegiance{
		Name: "Allegiance 1",
	})

	if len(allegiances) == 0 {
		return api.GetArmiesJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrAllegiancesNotFound})
	}

	return api.GetAllegiancesJSON200Response(allegiances)
}
