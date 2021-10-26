package session

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"travalite/internal/models"
	"travalite/pkg/httputils"
)

type Handlers struct {
	useCase UseCase
}

func NewHandlers(useCase UseCase) *Handlers {
	return &Handlers{
		useCase: useCase,
	}
}

func (h *Handlers) LogOut(w http.ResponseWriter, r *http.Request) {
	reqId := rand.Uint64()

	s := models.Session{}
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		httputils.Respond(w, r, reqId, http.StatusInternalServerError, err.Error())
		return
	}
	err := h.useCase.DelSession(s)
	if err != nil {
		httputils.Respond(w, r, reqId, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.Respond(w, r, reqId, http.StatusOK, nil)
}
