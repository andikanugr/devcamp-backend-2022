package main

import (
	"net/http"

	bookhandler "github.com/andikanugr/devcamp-backend-2022/clean-architecture/internal/handler/http/book"
	"github.com/go-chi/chi"
)

func newRoutes(bookHandler *bookhandler.Handler) *chi.Mux {
	router := chi.NewRouter()

	router.MethodFunc(http.MethodGet, "/v1/book", bookHandler.GetBooks)
	router.MethodFunc(http.MethodPost, "/v1/book", bookHandler.Create)

	return router
}
