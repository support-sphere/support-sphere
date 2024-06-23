package delivery

import (
	"log"
	"net/http"

	"github.com/support-sphere/support-sphere/config"
)

type HTTP struct {
	Config *config.Config
}

func LoggerInfoCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("INFO: %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// func (h *HTTP) setupSwaggerDocs() {
// 	if h.Config.Server.Env == "development" {
// 		docs.SwaggerInfo.Title = h.Config.App.Name
// 		docs.SwaggerInfo.Version = h.Config.App.Revision
// 		swaggerURL := fmt.Sprintf("%s/swagger/doc.json", h.Config.App.URL)
// 		h.mux.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(swaggerURL)))
// 		log.Info().Str("url", swaggerURL).Msg("Swagger documentation enabled.")
// 	}
// }
