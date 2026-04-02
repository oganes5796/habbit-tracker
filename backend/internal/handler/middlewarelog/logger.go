package middlewarelog

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/oganes5796/habbit-tracker/pkg/logger"
	"go.uber.org/zap"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(p []byte) (int, error) {
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	n, err := rw.ResponseWriter.Write(p)
	rw.bytes += n
	return n, err
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rid := middleware.GetReqID(r.Context())
		ctx := r.Context()
		if rid != "" {
			ctx = context.WithValue(ctx, logger.RequestIDKey, rid)
		}

		// 🟢 Лог старта запроса
		logger.Info(ctx, "request_started",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("query", r.URL.RawQuery),
			zap.String("remote_ip", r.RemoteAddr),
			zap.String("user_agent", r.UserAgent()),
		)

		rw := &responseWriter{ResponseWriter: w}
		next.ServeHTTP(rw, r.WithContext(ctx))

		// 🔴 Лог завершения запроса
		logger.Info(ctx, "request_finished",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Int("status", rw.status),
			zap.Int("bytes", rw.bytes),
			zap.Duration("elapsed", time.Since(start)),
			zap.String("remote_ip", r.RemoteAddr),
			zap.String("user_agent", r.UserAgent()),
		)
	})
}
