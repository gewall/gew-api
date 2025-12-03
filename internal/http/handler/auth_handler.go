package handler

import (
	"errors"
	"gew/internal/http/dto"
	"gew/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AuthHandler interface {
	AuthRoute(chi.Router)
	Ping(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	Register(http.ResponseWriter, *http.Request)
	Refresh(http.ResponseWriter, *http.Request)
}

type authHandler struct {
	svc service.AuthService
}

func NewAuth(svc service.AuthService) AuthHandler {
	return &authHandler{
		svc,
	}
}

func (h *authHandler) AuthRoute(r chi.Router) {
	r.Get("/ping", h.Ping)
	r.Post("/login", h.Login)
	r.Post("/register", h.Register)
	r.Get("/refresh", h.Refresh)
}

func (h *authHandler) Ping(w http.ResponseWriter, r *http.Request) {
	if err := h.svc.Ping(); err != nil {
		w.Write([]byte("Something wrong with auth service"))
	}
	w.Write([]byte("Auth handler is work well!"))
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userInput dto.LoginRequest
	if r.ContentLength == 0 {

		render.Render(w, r, dto.ErrInvalidRequest(errors.New("Field Empty")))
		return
	}

	if err := render.Bind(r, &userInput); err != nil {
		render.Render(w, r, dto.ErrInvalidRequest(err))
		return
	}

	res, err := h.svc.Login(&userInput)
	if err != nil {
		render.Render(w, r, dto.ErrInvalidRequest(err))
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    res["refreshToken"].(string),
		HttpOnly: true,
		MaxAge:   int(res["refreshTokenExp"].(float64)),
		Path:     "/",
	})

	render.Render(w, r, dto.NewResponse("Login Successfully", dto.LoginResponse{
		AccessToken: res["accessToken"].(string),
	}))
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userInput dto.RegisterRequest
	if err := render.Bind(r, &userInput); err != nil {
		render.Render(w, r, dto.ErrInvalidRequest(err))
		return
	}

	res, err := h.svc.Register(&userInput)
	if err != nil {
		render.Render(w, r, dto.ErrInvalidRequest(err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    res["refreshToken"].(string),
		HttpOnly: true,
		MaxAge:   int(res["refreshTokenExp"].(float64)),
		Path:     "/",
	})

	render.Render(w, r, dto.NewResponse("Register Successfully", dto.RegisterResponse{
		AccessToken: res["accessToken"].(string),
	}))

}

func (h *authHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	refrToken, err := r.Cookie("refreshToken")
	if err != nil {
		render.Render(w, r, dto.ErrInvalidRequest(err))
		return
	}

	res, err := h.svc.Refresh(refrToken.Value)
	if err != nil {
		render.Render(w, r, dto.ErrInvalidRequest(err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    res["refreshToken"].(string),
		HttpOnly: true,
		MaxAge:   int(res["refreshTokenExp"].(float64)),
		Path:     "/",
	})

	render.Render(w, r, dto.NewResponse("Refresh Successfully", dto.LoginResponse{
		AccessToken: res["accessToken"].(string),
	}))

}
