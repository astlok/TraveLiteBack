package profile

import "net/http"

type Handlers struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) *Handlers {
	return &Handlers{
		useCase: useCase,
	}
}

func (h *Handlers) AuthProfile(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) RegistrationProfile(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) ChangeProfile(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) GetProfile(w http.ResponseWriter, r *http.Request) {

}
