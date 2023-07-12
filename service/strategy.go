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

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event":       "GetStrategyByID",
		"strategy_id": id,
	}).Info()

	strategy, err := s.repo.GetGrandStrategyByID(r.Context(), id)
	if err != nil {
		return api.GetGrandStrategyByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrGrandStrategyNotFound})
	}

	return api.GetGrandStrategyByIDJSON200Response(api.GrandStrategy(*strategy))
}

func (s *WarhammerService) GetGrandStrategies(w http.ResponseWriter, r *http.Request) *api.Response {

	logging.NewLogrus(r.Context()).WithFields(logrus.Fields{
		"event": "GetStrategies",
	}).Info()

	strategies, err := s.repo.GetGrandStrategies(r.Context())
	if err != nil {
		return api.GetArmiesJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrGrandStrategiesNotFound})
	}

	if len(strategies) == 0 {
		return api.GetArmiesJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrGrandStrategiesNotFound})
	}

	return api.GetGrandStrategiesJSON200Response(strategies)
}
