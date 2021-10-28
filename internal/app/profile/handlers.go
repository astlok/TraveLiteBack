package profile

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
	"travalite/internal/app/session"
	"travalite/internal/models"
	"travalite/pkg/constants"
	customErrors "travalite/pkg/errors"
	"travalite/pkg/httputils"
)

type Handlers struct {
	profileUseCase UseCase
	sessionUseCase session.UseCase
}

func NewHandler(profileUseCase UseCase, sessionUseCase session.UseCase) *Handlers {
	return &Handlers{
		profileUseCase: profileUseCase,
		sessionUseCase: sessionUseCase,
	}
}

func (h *Handlers) AuthProfile(w http.ResponseWriter, r *http.Request) {
	reqId := rand.Uint64()
	u := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		httputils.Respond(w, r, reqId, http.StatusInternalServerError, err.Error())
		return
	}

	u, err := h.profileUseCase.Auth(u)
	if err != nil {
		if errors.Is(err, customErrors.BadAuth) {
			httputils.Respond(w, r, reqId, http.StatusUnauthorized, err.Error())
			return
		}
		httputils.Respond(w, r, reqId, http.StatusInternalServerError, err.Error())
		return
	}
	s, err := h.sessionUseCase.Create(u)
	u.AuthToken = s.AuthToken

	httputils.Respond(w, r, reqId, http.StatusOK, u)
}

func (h *Handlers) RegistrationProfile(w http.ResponseWriter, r *http.Request) {
	reqId := rand.Uint64()
	u := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		httputils.Respond(w, r, reqId, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.profileUseCase.Create(*u)
	if err != nil {
		if errors.Is(err, customErrors.DuplicateEmail) || errors.Is(err, customErrors.DuplicateNickName) {
			httputils.Respond(w, r, reqId, http.StatusConflict, err.Error())
			return
		}
		httputils.Respond(w, r, reqId, http.StatusInternalServerError, err.Error())
		return
	}
	u.ID = id

	s, err := h.sessionUseCase.Create(*u)
	if err != nil {
		httputils.Respond(w, r, reqId, http.StatusInternalServerError, err.Error())
	}

	httputils.Respond(w, r, reqId, http.StatusCreated, s)
}

func (h *Handlers) ChangeProfile(w http.ResponseWriter, r *http.Request) {
	reqId := rand.Uint64()

	u := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		if errors.Is(err, customErrors.DuplicateEmail) || errors.Is(err, customErrors.DuplicateNickName) {
			httputils.Respond(w, r, reqId, http.StatusConflict, err.Error())
			return
		}
		httputils.Respond(w, r, reqId, http.StatusInternalServerError, err.Error())
		return
	}

	u.ID = r.Context().Value(constants.CtxUserID).(uint64)
	err := h.profileUseCase.ChangeProfile(*u)
	if err != nil {
		httputils.Respond(w, r, reqId, http.StatusInternalServerError, err.Error())
		return
	}
	httputils.Respond(w, r, reqId, http.StatusOK, nil)
}

func (h *Handlers) GetProfile(w http.ResponseWriter, r *http.Request) {
	reqId := rand.Uint64()

	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		httputils.Respond(w, r, reqId, http.StatusInternalServerError, err.Error())
		return
	}

	u, err := h.profileUseCase.GetUser(id)
	if err != nil {
		if errors.Is(err, customErrors.UserNotFound) {
			httputils.Respond(w, r, reqId, http.StatusNotFound, err.Error())
			return
		}
		httputils.Respond(w, r, reqId, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.Respond(w, r, reqId, http.StatusOK, u)
}
