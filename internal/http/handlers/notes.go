package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/MatveySotnikov/notes-api/internal/core"
	"github.com/MatveySotnikov/notes-api/internal/core/service"
)

type Handler struct {
	Service service.NoteService
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var n core.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	created, err := h.Service.Create(r.Context(), n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(created)
}

func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	n, err := h.Service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(n)
}

func (h *Handler) ListNotes(w http.ResponseWriter, r *http.Request) {
	notes, _ := h.Service.List(r.Context())
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(notes)
}

func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	var n core.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	updated, err := h.Service.Update(r.Context(), id, n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(updated)
}

func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err := h.Service.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
