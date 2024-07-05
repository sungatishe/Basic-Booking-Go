package main

import (
	"net/http"

	"github.com/sungatishe/pkgs/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func route() http.Handler {
	fileServer := http.FileServer(http.Dir("../static/"))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/majors", handlers.Repo.Majors)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/make-reservations", handlers.Repo.MakeReservations)
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostSearchAvailability)
	mux.Get("/contacts", handlers.Repo.Contacts)

	return mux
}
