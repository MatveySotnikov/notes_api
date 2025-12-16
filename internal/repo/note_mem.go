package repo

import (
	"errors"
	"sync"
	"time"

	"github.com/MatveySotnikov/notes-api/internal/core"
)

var ErrNoteNotFound = errors.New("note not found")

type NoteRepoMem struct {
	mu    sync.RWMutex
	notes map[int64]*core.Note
	next  int64
}

func NewNoteRepoMem() *NoteRepoMem {
	return &NoteRepoMem{
		notes: make(map[int64]*core.Note),
	}
}

func (r *NoteRepoMem) Create(n core.Note) (core.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.next++
	n.ID = r.next
	n.CreatedAt = time.Now().UTC()

	r.notes[n.ID] = &n
	return n, nil
}

func (r *NoteRepoMem) GetByID(id int64) (core.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	n, ok := r.notes[id]
	if !ok {
		return core.Note{}, ErrNoteNotFound
	}
	return *n, nil
}

func (r *NoteRepoMem) List() ([]core.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]core.Note, 0, len(r.notes))
	for _, n := range r.notes {
		result = append(result, *n)
	}
	return result, nil
}

func (r *NoteRepoMem) Update(id int64, n core.Note) (core.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, ok := r.notes[id]
	if !ok {
		return core.Note{}, ErrNoteNotFound
	}

	existing.Title = n.Title
	existing.Content = n.Content
	now := time.Now().UTC()
	existing.UpdatedAt = &now

	return *existing, nil
}

func (r *NoteRepoMem) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.notes[id]; !ok {
		return ErrNoteNotFound
	}
	delete(r.notes, id)
	return nil
}
