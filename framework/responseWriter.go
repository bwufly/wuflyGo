package WuflyGo

import "net/http"

// responseWriter implement http.ResponseWriter
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
}

