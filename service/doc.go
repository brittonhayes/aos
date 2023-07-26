package service

import (
	"net/http"

	"github.com/brittonhayes/warhammer"
	"github.com/brittonhayes/warhammer/api"
)

func (s *WarhammerService) GetDocs(w http.ResponseWriter, r *http.Request) *api.Response {
	resp := api.Response{}
	resp.ContentType("text/html")
	resp.Status(http.StatusOK)

	w.Write(warhammer.OPEN_API_DOCS)

	return &resp
}
