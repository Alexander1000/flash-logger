package auth

import (
	"net/http"

	"flash-logger/response/json"
)

type Handler struct {
	fallback http.Handler
}

func NewAuthHandler(fallback http.Handler) http.Handler {
	return &Handler{fallback:fallback}
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	authHeader := req.Header.Get("Authorization")
	defer req.Body.Close()
	if authHeader == "" {
		json.Reply(resp, json.ErrorUnauthorized, http.StatusUnauthorized)
		return
	}
	h.fallback.ServeHTTP(resp, req)
}
