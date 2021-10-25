package httputils

import (
	"encoding/json"
	"net/http"
	"travalite/pkg/logger"
)

func Respond(w http.ResponseWriter, r *http.Request, requestId uint64, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			logger.LoggingError(requestId, err)
			return
		}
	}
	logger.LoggingResponse(requestId, code)
}
