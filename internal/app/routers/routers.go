package routers

import (
	"github.com/SamoylenkoVadim/golang-practicum/internal/app/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter(h *handlers.Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", h.PostHandler)
	router.Get("/{id}", h.GetHandler)

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad request", http.StatusBadRequest)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad request", http.StatusBadRequest)
	})

	return router
}
