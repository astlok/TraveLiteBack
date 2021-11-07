package region

import (
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
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

func (h *Handlers) GetRegionInfo(w http.ResponseWriter, r *http.Request) {
	reqID := rand.Uint64()

	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)

	region, err := h.useCase.GetRegionInfo(id)
	if err != nil {
		httputils.Respond(w, r , reqID, http.StatusInternalServerError, err)
	}

	httputils.Respond(w, r, reqID, http.StatusOK, region)
}

func (h *Handlers) GetRegions(w http.ResponseWriter, r *http.Request) {
	reqID := rand.Uint64()
	regions, err := h.useCase.GetRegions()
	if err != nil {
		httputils.Respond(w, r , reqID, http.StatusInternalServerError, err)
	}
	httputils.Respond(w, r, reqID, http.StatusOK, regions)
}