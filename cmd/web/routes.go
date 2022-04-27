package main

import (
	"net/http"

	"github.com/AmanSrivastava2000/bookings/cmd/pkg/config"
	"github.com/AmanSrivastava2000/bookings/cmd/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	//setting up routes:
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//getting data from static folder.
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
