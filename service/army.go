package service

import (
	"encoding/json"
	"net/http"

	"github.com/brittonhayes/warhammer/api"
	"github.com/brittonhayes/warhammer/internal/logging"
	"github.com/sirupsen/logrus"
)

const (
	ErrArmyNotFound   = "Army not found with this ID"
	ErrArmiesNotFound = "No armies found"
)

func (s *WarhammerService) GetArmyByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event":   "GetArmyByID",
		"army_id": id,
	}).Info()

	army, err := s.repo.GetArmyByID(r.Context(), id)
	if err != nil {
		return api.GetArmyByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrArmyNotFound})
	}

	return api.GetArmyByIDJSON200Response(*army)
}

func (s *WarhammerService) GetArmies(w http.ResponseWriter, r *http.Request) *api.Response {

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event": "GetArmies",
	}).Info()

	armies, err := s.repo.GetArmies(r.Context())
	if err != nil {
		return api.GetArmiesJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrArmiesNotFound})
	}

	if len(armies) == 0 {
		return api.GetArmiesJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrArmiesNotFound})
	}

	return api.GetArmiesJSON200Response(armies)
}

func (s *WarhammerService) UpdateArmyByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event":   "UpdateArmyByID",
		"army_id": id,
	}).Info()

	army := new(api.Army)
	err := json.NewDecoder(r.Body).Decode(&army)
	if err != nil {
		return api.UpdateArmyByIDJSON404Response(api.Error{Code: http.StatusBadRequest, Message: err.Error()})
	}

	army, err = s.repo.UpdateArmyByID(r.Context(), id, army)
	if err != nil {
		return api.UpdateArmyByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrArmyNotFound})
	}

	return api.UpdateArmyByIDJSON200Response(*army)
}
