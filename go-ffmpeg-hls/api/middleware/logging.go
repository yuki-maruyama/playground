package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type loggingWriter struct {
	http.ResponseWriter
	code int
}

func newLoggingWriter(w http.ResponseWriter) *loggingWriter {
	return &loggingWriter{ResponseWriter: w, code: http.StatusInternalServerError}
}

func (lw *loggingWriter) WriteHeader(code int) {
	lw.code = code
	lw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}
		reqTime := time.Now()
		rlw := newLoggingWriter(w)

		next.ServeHTTP(rlw, r)
		slog.InfoContext(r.Context(), "response", "remote_host", r.RemoteAddr, "method", r.Method, "uri", r.RequestURI, "status", rlw.code, "duration", time.Since(reqTime).Seconds())
	}
	return http.HandlerFunc(fn)
}
