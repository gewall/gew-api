package main

import (
	"gew/internal/config"
	"gew/internal/db"
	"gew/internal/http/handler"

	"gew/internal/repository"
	"gew/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

type Response struct {
	Message string
	Data    any
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func main() {
	err := config.LoadEnv()

	if err != nil {
		log.Fatalf("Something wrong with the env: %s", err.Error())
	}
	postgresqlDb := db.Init()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.Render(w, r, &Response{
			Message: "Server is doing right",
			Data:    nil,
		})
	})

	// user route setup
	userRepo := repository.NewUser(postgresqlDb)
	tokenRepo := repository.NewToken(postgresqlDb)
	linkRepo := repository.NewLink(postgresqlDb)
	// auth route setup
	authHdr := handler.NewAuth(service.NewAuth(userRepo, tokenRepo))
	// link route setup
	linkhdr := handler.NewLink(service.NewLink(linkRepo))

	// public route
	r.Group(func(r chi.Router) {
		r.Route("/auth", authHdr.AuthRoute)
		r.Route("/link", linkhdr.PublicLinkRoute)
	})
	// private route
	r.Route("/v1", func(r chi.Router) {
		tokenAuth := config.JwtMiddleware()
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Route("/link", linkhdr.PrivateLinkRoute)
	})

	http.ListenAndServe(":8080", r)
}
