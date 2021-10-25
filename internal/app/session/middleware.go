package session

import (
	"net/http"
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

		next.ServeHTTP(w, r)
	})
}
