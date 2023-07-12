package service

import (
	"net/http"

	"github.com/brittonhayes/warhammer/api"
	"github.com/brittonhayes/warhammer/internal/logging"
	"github.com/sirupsen/logrus"
)

const (
	ErrGrandStrategyNotFound   = "Strategy not found with this ID"
	ErrGrandStrategiesNotFound = "No strategys found"
)

func (s *WarhammerService) GetGrandStrategyByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {
	var strategy api.GrandStrategy
	strategy.ID = &id

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event":       "GetStrategyByID",
		"strategy_id": id,
	}).Info()

	if strategy.ID == nil {
		return api.GetGrandStrategyByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrGrandStrategyNotFound})
	}

	return api.GetGrandStrategyByIDJSON200Response(strategy)
}

func (s *WarhammerService) GetGrandStrategies(w http.ResponseWriter, r *http.Request) *api.Response {
	var strategys []api.GrandStrategy

	strategys = append(strategys, api.GrandStrategy{
		Name: "Strategy 1",
	})

	if len(strategys) == 0 {
		return api.GetArmiesJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrGrandStrategiesNotFound})
	}

	return api.GetGrandStrategiesJSON200Response(strategys)
}
