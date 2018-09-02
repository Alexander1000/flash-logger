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
	jsonResponse.Reply(resp, response{Result: true}, http.StatusOK)
}
