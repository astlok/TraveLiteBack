package trek

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
	"travalite/internal/models"
	"travalite/pkg/constants"
	"travalite/pkg/httputils"
)

type Handlers struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) *Handlers {
	return &Handlers{
		useCase: useCase,
	}
}

func (h *Handlers) CreateTrek(w http.ResponseWriter, r *http.Request) {
	reqID := rand.Uint64()

	ID := r.Context().Value(constants.CtxUserID).(uint64)

	var t models.Trek

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusBadRequest, err)
		return
	}


	t, err = h.useCase.CreateTrek(ID, t)

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusConflict, err)
		return
	}

	httputils.Respond(w, r, reqID, http.StatusCreated, t)
}

func (h *Handlers) GetTreks(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) GetTrekInfo(w http.ResponseWriter, r *http.Request) {
	reqID := rand.Uint64()

	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusInternalServerError, err)
		return
	}

	trek, err := h.useCase.GetTrekInfo(ID)

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusConflict, err)
		return
	}

	if trek.ID == 0 {
		httputils.Respond(w, r, reqID, http.StatusNotFound, "{}")
		return
	}

	httputils.Respond(w, r, reqID, http.StatusOK, trek)
}

func (h *Handlers) ChangeTrek(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) DelTrek(w http.ResponseWriter, r *http.Request) {
	reqID := rand.Uint64()

	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusInternalServerError, err)
		return
	}

	err = h.useCase.DeleteTrek(ID)

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusConflict, err)
	}

	httputils.Respond(w, r, reqID, http.StatusOK, "{}")
}

func (h *Handlers) GetUsersTreks(w http.ResponseWriter, r *http.Request) {
	reqID := rand.Uint64()

	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusInternalServerError, err)
		return
	}

	treks, err := h.useCase.GetUsersTreks(ID)
	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusConflict, err)
		return
	}

	httputils.Respond(w, r, reqID, http.StatusOK, treks)
}

func (h *Handlers) SearchTrek(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) CreatComment(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) GetTrekComments(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) RateTrek(w http.ResponseWriter, r *http.Request) {

}

