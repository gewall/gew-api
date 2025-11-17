package handler

import (
	"gew/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type authHandler struct {
	svc service.AuthService
}

func NewAuth(svc service.AuthService) *authHandler {
	return &authHandler{
		svc,
	}
}

func (h *authHandler) AuthRoute(r chi.Router) {
	r.Get("/ping", h.Ping)
}

func (h *authHandler) Ping(w http.ResponseWriter, r *http.Request) {
	if err := h.svc.Ping(); err != nil {
		w.Write([]byte("Something wrong with auth service"))
	}
	w.Write([]byte("Auth handler is work well!"))
}
