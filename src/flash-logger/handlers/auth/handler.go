package auth

import (
	"net/http"

	"flash-logger/response/json"
	"flash-logger/config"
)

type Handler struct {
	fallback http.Handler
	projects []config.Project
}

func NewAuthHandler(fallback http.Handler, projects []config.Project) http.Handler {
	return &Handler{
		fallback:fallback,
		projects: projects,
	}
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	authHeader := req.Header.Get("Authorization")
	defer req.Body.Close()
	if authHeader == "" {
		json.Reply(resp, json.ErrorUnauthorized, http.StatusUnauthorized)
		return
	}

	if len(authHeader) < 10 {
		json.Reply(resp, json.ErrorNotImplemented, http.StatusNotImplemented)
		return
	}

	// Authorization: Bearer <token>
	if authType := authHeader[0:7]; authType != "Bearer " {
		json.Reply(resp, json.ErrorNotImplemented, http.StatusNotImplemented)
		return
	}

	h.fallback.ServeHTTP(resp, req)
}
