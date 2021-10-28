package session

import (
	"context"
	"math/rand"
	"net/http"
	"travalite/pkg/constants"
	"travalite/pkg/httputils"
)

type Middleware struct {
	useCase UseCase
}

func NewMiddleware(useCase UseCase) *Middleware {
	return &Middleware{
		useCase: useCase,
	}
}

func (m *Middleware) CheckSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqId := rand.Uint64()

		sessionID := r.Header.Get("X-Auth-token")

		u, err := m.useCase.Check(sessionID)
		if err != nil {
			httputils.Respond(w, r, reqId, http.StatusForbidden, "bad auth-token")
			return
		}

		ctx := context.WithValue(r.Context(), constants.CtxUserID, u.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
