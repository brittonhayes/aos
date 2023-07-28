package service

import (
	"net/http"

	"github.com/brittonhayes/aos/api"
	"github.com/brittonhayes/aos/internal/logging"
	"github.com/sirupsen/logrus"
)

const (
	ErrAllianceNotFound  = "Alliance not found with this ID"
	ErrAlliancesNotFound = "No alliances found"
)

func (s *Service) GetGrandAllianceByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event":       "GetGrandAllianceByID",
		"alliance_id": id,
	}).Info()

	alliance, err := s.repo.GetGrandAllianceByID(r.Context(), id)
	if err != nil {
		return api.GetGrandAllianceByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrAllianceNotFound})
	}

	return api.GetGrandAllianceByIDJSON200Response(*alliance)
}

func (s *Service) GetGrandAlliances(w http.ResponseWriter, r *http.Request) *api.Response {

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event": "GetGrandAlliances",
	}).Info()

	alliances, err := s.repo.GetGrandAlliances(r.Context())
	if err != nil {
		return api.GetGrandAlliancesJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrAlliancesNotFound})
	}

	if len(alliances) == 0 {
		return api.GetGrandAlliancesJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrAlliancesNotFound})
	}

	return api.GetGrandAlliancesJSON200Response(alliances)
}
