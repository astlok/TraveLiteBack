package session

import "net/http"

type Handlers struct {
	useCase UseCase
}

func NewHandlers(useCase UseCase) *Handlers {
	return &Handlers{
		useCase: useCase,
	}
}

func (h *Handlers) LogOut(w http.ResponseWriter, r *http.Request) {

}
