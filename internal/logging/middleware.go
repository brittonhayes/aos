package logging

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NewLogrus(r.Context()).WithFields(logrus.Fields{
			"url": r.URL.String(),
		}).Info(r.Method)

		h.ServeHTTP(w, r)
	})
}
