package auth

import "net/http"

type Handler struct {
	fallback http.Handler
}

func NewAuthHandler(fallback http.Handler) http.Handler {
	return &Handler{fallback:fallback}
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	h.fallback.ServeHTTP(resp, req)
}
