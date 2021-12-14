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

	t.UserID = ID

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
	reqID := rand.Uint64()

	treks, err := h.useCase.GetTreks()

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusConflict, err)
		return
	}

	httputils.Respond(w, r, reqID, http.StatusOK, treks)
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
	reqID := rand.Uint64()
	ID := r.Context().Value(constants.CtxUserID).(uint64)
	t := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusInternalServerError, err)
		return
	}

	trek, err := h.useCase.ChangeTreks(ID, t)

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusConflict, err)
		return
	}

	httputils.Respond(w, r, reqID, http.StatusOK, trek)
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
	reqID := rand.Uint64()
	ID := r.Context().Value(constants.CtxUserID).(uint64)

	var c models.TrekComment
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusInternalServerError, err)
		return
	}

	comment, err := h.useCase.CreateComment(ID, c)
	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusConflict, err)
		return
	}

	httputils.Respond(w, r, reqID, http.StatusCreated, comment)
}

func (h *Handlers) GetTrekComments(w http.ResponseWriter, r *http.Request) {
	reqID := rand.Uint64()

	comments, err := h.useCase.GetComments()
	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusConflict, err)
		return
	}

	httputils.Respond(w, r, reqID, http.StatusOK, comments)
}

func (h *Handlers) RateTrek(w http.ResponseWriter, r *http.Request) {
	reqID := rand.Uint64()
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusInternalServerError, err)
		return
	}

	trekRate := make(map[string]uint64)
	err = json.NewDecoder(r.Body).Decode(&trekRate)

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusInternalServerError, err)
		return
	}

	err = h.useCase.RateTrek(ID, trekRate)

	if err != nil {
		httputils.Respond(w, r, reqID, http.StatusConflict, err)
		return
	}

	httputils.Respond(w, r, reqID, http.StatusOK, "{}")
}

