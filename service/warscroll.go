package service

import (
	"net/http"

	"github.com/brittonhayes/warhammer/api"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	ErrWarscrollNotFound  = "Warscroll not found with this ID"
	ErrWarscrollsNotFound = "No warscrolls found"
)

func (s *WarhammerService) GetWarscrollByID(w http.ResponseWriter, r *http.Request, id string) *api.Response {
	span := trace.SpanFromContext(r.Context())
	defer span.End()

	warscroll, err := s.repo.GetWarscrollByID(r.Context(), id)
	if err != nil {
		return api.GetWarscrollByIDJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrWarscrollNotFound})
	}
	span.SetAttributes(attribute.String("warscroll.allegiance_id", *warscroll.AllegianceID))
	span.SetAttributes(attribute.String("warscroll.grand_alliance_id", *warscroll.GrandAllianceID))

	warscroll.Allegiance, err = s.repo.GetAllegianceByID(r.Context(), *warscroll.AllegianceID)
	if err != nil {
		span.SetAttributes(semconv.ExceptionMessageKey.String(err.Error()))
		span.AddEvent("failed to find allegiance")
	}

	warscroll.GrandAlliance, err = s.repo.GetGrandAllianceByID(r.Context(), *warscroll.GrandAllianceID)
	if err != nil {
		span.SetAttributes(semconv.ExceptionMessageKey.String(err.Error()))
		span.AddEvent("failed to find grand alliance")
	}

	return api.GetWarscrollByIDJSON200Response(*warscroll)
}

func (s *WarhammerService) GetWarscrolls(w http.ResponseWriter, r *http.Request, params api.GetWarscrollsParams) *api.Response {
	span := trace.SpanFromContext(r.Context())
	defer span.End()

	response, err := s.repo.GetWarscrolls(r.Context(), params)
	if err != nil {
		return api.GetWarscrollsJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrWarscrollsNotFound})
	}

	if len(response) == 0 {
		return api.GetWarscrollsJSON404Response(api.Error{Code: http.StatusNotFound, Message: ErrWarscrollsNotFound})
	}

	return api.GetWarscrollsJSON200Response(response)
}
