package logs

import (
	"net/http"

	jsonResponse "flash-logger/response/json"
	"flash-logger/storage"
)

type Handler struct {
	storage storage.Repository
}

func New(storage storage.Repository) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	// @todo забирать из входных параметров
	messages := h.storage.GetLastMessages(1, 20, 0)
	jsonResponse.Reply(resp, response{Result: messages}, http.StatusOK)
}
