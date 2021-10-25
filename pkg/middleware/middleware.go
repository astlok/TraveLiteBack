package middleware

import (
	"context"
	"math/rand"
	"net/http"
	"travalite/pkg/logger"
	"travalite/pkg/types"
)

func LoggingRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := rand.Uint64()
		logger.LoggingRequest(reqID, r.URL, r.Method)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), constants.ReqID, reqID)))
	})
}