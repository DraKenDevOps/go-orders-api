package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/ok", Ok)

	return router
}

func Ok(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
}
