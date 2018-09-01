package event

import (
	"net/http"

	"flash-logger/response/json"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	json.Reply(resp, response{Result: true}, http.StatusOK)
}
