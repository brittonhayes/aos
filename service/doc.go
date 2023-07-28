package service

import (
	"net/http"

	"github.com/brittonhayes/aos"
	"github.com/brittonhayes/aos/api"
)

func (s *aosService) GetDocs(w http.ResponseWriter, r *http.Request) *api.Response {
	resp := api.Response{}
	resp.ContentType("text/html")
	resp.Status(http.StatusOK)

	w.Write(aos.OPEN_API_DOCS)

	return &resp
}
