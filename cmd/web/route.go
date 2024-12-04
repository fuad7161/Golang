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
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)

	fileServer := http.FileServer(http.Dir("../../static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
