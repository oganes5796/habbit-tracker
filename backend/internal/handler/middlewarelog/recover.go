package middlewarelog

import (
	"net/http"
	"runtime/debug"

	"github.com/oganes5796/habbit-tracker/pkg/logger"
	"go.uber.org/zap"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				ctx := r.Context()

				logger.Error(ctx, "panic recovered",
					zap.Any("panic_value", rvr),
					zap.String("stack", string(debug.Stack())),
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.String("remote_ip", r.RemoteAddr),
				)

				// Не забываем ответить клиенту
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"error":"internal server error"}`))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
