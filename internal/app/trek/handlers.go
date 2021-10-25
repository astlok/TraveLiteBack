package trek

import "net/http"

type Handlers struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) *Handlers {
	return &Handlers{
		useCase: useCase,
	}
}

func (h *Handlers) CreateTrek(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) GetTreks(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) GetTrekInfo(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) ChangeTrek(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) DelTrek(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) GetUsersTreks(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) SearchTrek(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) CreatComment(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) GetTrekComments(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) RateTrek(w http.ResponseWriter, r *http.Request) {

}

