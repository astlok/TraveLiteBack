package region

import "net/http"

type Handlers struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) *Handlers {
	return &Handlers{
		useCase: useCase,
	}
}

func (h *Handlers) GetRegionInfo(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) GetRegions(w http.ResponseWriter, r *http.Request) {

}