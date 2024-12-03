package main

import (
	"github.com/fuad7161/Golang/tree/Project/Bookings/pkg/config"
	"github.com/fuad7161/Golang/tree/Project/Bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("../../static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
