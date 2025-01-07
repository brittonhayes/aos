package sunset

import (
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/brittonhayes/aos"
	"github.com/brittonhayes/aos/internal/logging"
)

// sunsetDate defines when the public API instance will be sunset
var sunsetDate = time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

// Middleware is a middleware that displays a sunset notice if the current date is after the sunset date
func Middleware(next http.Handler) http.Handler {
	tmpl := template.Must(template.New("sunset").Parse(string(aos.HTML_SUNSET)))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if time.Now().After(sunsetDate) && strings.Contains(r.Host, "aos-api.com") {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			data := map[string]interface{}{
				"SunsetDate": sunsetDate.Format("January 2, 2006"),
			}
			if err := tmpl.Execute(w, data); err != nil {
				logging.NewLogrus(r.Context()).Errorf("Error executing sunset template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
