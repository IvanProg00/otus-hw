package internalhttp

import (
	"fmt"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{w, http.StatusOK}
		next.ServeHTTP(lrw, r)
		end := time.Since(start)
		s.Logger.Info(fmt.Sprintf("%s %s %s %s %d %d \"%s\"", r.RemoteAddr, r.Method, r.URL, r.Proto, lrw.statusCode, end.Milliseconds(), r.UserAgent()))

	})
}
