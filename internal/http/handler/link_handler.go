package handler

import (
	"fmt"
	"gew/internal/http/dto"
	"gew/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type LinkHandler interface {
	PublicLinkRoute(chi.Router)
	PrivateLinkRoute(chi.Router)
	FindBySlug(http.ResponseWriter, *http.Request)
	CreateLink(http.ResponseWriter, *http.Request)
	FindsById(http.ResponseWriter, *http.Request)
	DeleteById(http.ResponseWriter, *http.Request)
}

type linkHandler struct {
	svc service.LinkService
}

func NewLink(svc service.LinkService) LinkHandler {
	return &linkHandler{svc}
}

func (h *linkHandler) PublicLinkRoute(r chi.Router) {
	r.Get("/{slug}", h.FindBySlug)
}

func (h *linkHandler) PrivateLinkRoute(r chi.Router) {
	// r.Get("/", h.FindBySlug)
	r.Post("/", h.CreateLink)
	r.Get("/{id}", h.FindsById)
	r.Delete("/{id}", h.DeleteById)
}

func (h *linkHandler) FindBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	res, err := h.svc.FindBySlug(slug)
	if err != nil {
		render.Render(w, r, dto.ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, dto.NewResponse("Link Found.", map[string]string{
		"destination": res["destination"].(string),
	}))
}

func (h *linkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	var linkInput dto.LinkRequest
	if err := render.Bind(r, &linkInput); err != nil {
		render.Render(w, r, dto.ErrInvalidRequest(err))
		return
	}

	fmt.Print(linkInput)

	if err := h.svc.CreateLink(&linkInput); err != nil {
		render.Render(w, r, dto.ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, dto.NewResponse("Create Link Successfully", map[string]string{}))

}

func (h *linkHandler) FindsById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.svc.FindsById(id)
	if err != nil {
		render.Render(w, r, dto.ErrInvalidRequest(err))
		return
	}

	fmt.Println(res)

	render.Render(w, r, dto.NewResponse("Finds Link Successfully", res))
}

func (h *linkHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.svc.DeleteById(id); err != nil {
		render.Render(w, r, dto.ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, dto.NewResponse("Link Deleted", nil))
}
