package handlerloggerwrapper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	lz "github.com/karincake/apem/logger-zerolog"
	lo "github.com/karincake/apem/loggero"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func SetLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		lz.O.Info().
			String("scope", "request").
			Int("status", wrapped.statusCode).
			String("method", r.Method).
			String("path", r.URL.Path).
			String("query", r.URL.RawQuery).
			String("duration", time.Since(start).String()).Send()
	})
}

func WriteJson(data any) {
	lo.I.Println("Showing additional info for the payload")
	js, err := json.Marshal(data)
	if err == nil {
		fmt.Println(string(js))
	} else {
		fmt.Println("error converting data or result to json:", err)
	}
}
