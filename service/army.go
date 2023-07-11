package service

import (
	"net/http"

	"github.com/brittonhayes/warhammer/api"
	"github.com/brittonhayes/warhammer/internal/logging"
	"github.com/sirupsen/logrus"
)

func (s *WarhammerService) GetArmyByID(w http.ResponseWriter, r *http.Request, id int64) *api.Response {
	var army api.Army
	army.ID = &id

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event":   "GetArmyByID",
		"army_id": id,
	}).Info()

	if army.ID == nil {
		return api.GetArmyByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrArmyNotFound})
	}

	return api.GetArmyByIDJSON200Response(army)
}

func (s *WarhammerService) GetArmies(w http.ResponseWriter, r *http.Request) *api.Response {
	var armies []api.Army

	armies = append(armies, api.Army{
		Name: "Army 1",
	})

	if len(armies) == 0 {
		return api.GetArmiesJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrArmiesNotFound})
	}

	return api.GetArmiesJSON200Response(armies)
}

func (s *WarhammerService) UpdateArmyByID(w http.ResponseWriter, r *http.Request, id int64) *api.Response {
	var army api.Army
	army.ID = &id

	return api.UpdateArmyByIDJSON200Response(army)
}
