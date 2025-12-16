package httpx

import (
	"github.com/go-chi/chi/v5"

	"github.com/MatveySotnikov/notes-api/internal/http/handlers"
)

func NewRouter(h *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1/notes", func(r chi.Router) {
		r.Post("/", h.CreateNote)
		r.Get("/", h.ListNotes)
		r.Get("/{id}", h.GetNote)
		r.Put("/{id}", h.UpdateNote)
		r.Delete("/{id}", h.DeleteNote)
	})

	return r
}
