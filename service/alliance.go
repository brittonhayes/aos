package service

import (
	"net/http"

	"github.com/brittonhayes/warhammer/api"
	"github.com/brittonhayes/warhammer/internal/logging"
	"github.com/sirupsen/logrus"
)

const (
	ErrAllianceNotFound  = "Alliance not found with this ID"
	ErrAlliancesNotFound = "No alliances found"
)

func (s *WarhammerService) GetGrandAllianceByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {
	var alliance api.GrandAlliance
	alliance.ID = &id

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event":       "GetGrandAllianceByID",
		"alliance_id": id,
	}).Info()

	if alliance.ID == nil {
		return api.GetGrandAllianceByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrAllianceNotFound})
	}

	return api.GetGrandAllianceByIDJSON200Response(alliance)
}

func (s *WarhammerService) GetGrandAlliances(w http.ResponseWriter, r *http.Request) *api.Response {
	var alliances []api.GrandAlliance

	alliances = append(alliances, api.GrandAlliance{
		Name: "Alliance 1",
	})

	if len(alliances) == 0 {
		return api.GetArmiesJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrAlliancesNotFound})
	}

	return api.GetGrandAlliancesJSON200Response(alliances)
}
