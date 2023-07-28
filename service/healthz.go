package service

import (
	"net/http"

	"github.com/brittonhayes/aos/api"
)

func (s *aosService) GetHealthz(w http.ResponseWriter, r *http.Request) *api.Response {
	resp := api.Health{Status: "ok"}
	return api.GetHealthzJSON200Response(resp)
}
