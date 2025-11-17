package main

import (
	"gew/internal/config"
	"gew/internal/db"
	"gew/internal/handler"
	"gew/internal/repository"
	"gew/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	// user Route
	userRepo := repository.NewUser(postgresqlDb)
	// auth Route
	authSvc := service.NewAuth(userRepo)
	authHdr := handler.NewAuth(authSvc)
	r.Route("/auth", authHdr.AuthRoute)

	http.ListenAndServe(":8080", r)
}
