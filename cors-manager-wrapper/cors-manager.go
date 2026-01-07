package corsmanagermw

import (
	"net/http"

	a "github.com/karincake/apem"
)

type CorsCfg struct {
	AllowedOrigins []string `yaml:"allowedOrigins"`
	AllowedMethod  string   `yaml:"allowedMethod"`
}

var cfg CorsCfg

func SetCors(next http.Handler) http.Handler {
	a.ParseSingleCfg(&cfg)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, origin := range cfg.AllowedOrigins {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", cfg.AllowedMethod)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
