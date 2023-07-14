package service

import (
	"net/http"

	"github.com/brittonhayes/warhammer/api"
	"github.com/brittonhayes/warhammer/web"
)

func (s *WarhammerService) GetDocs(w http.ResponseWriter, r *http.Request) *api.Response {
	resp := api.Response{}
	resp.ContentType("text/html")
	resp.Status(http.StatusOK)

	w.Write(web.OPEN_API_DOCS)

	return &resp
}
