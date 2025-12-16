package service

import (
	"context"
	"errors"
	"strings"

	"github.com/MatveySotnikov/notes-api/internal/core"
	"github.com/MatveySotnikov/notes-api/internal/repo"
)

var (
	ErrInvalidNote  = errors.New("invalid note")
	ErrNoteNotFound = errors.New("note not found")
)

type NoteService interface {
	Create(ctx context.Context, note core.Note) (core.Note, error)
	GetByID(ctx context.Context, id int64) (core.Note, error)
	List(ctx context.Context) ([]core.Note, error)
	Update(ctx context.Context, id int64, note core.Note) (core.Note, error)
	Delete(ctx context.Context, id int64) error
}

type noteService struct {
	repo repo.NoteRepository
}

func NewNoteService(r repo.NoteRepository) NoteService {
	return &noteService{repo: r}
}

func (s *noteService) Create(ctx context.Context, n core.Note) (core.Note, error) {
	if strings.TrimSpace(n.Title) == "" {
		return core.Note{}, ErrInvalidNote
	}
	return s.repo.Create(n)
}

func (s *noteService) GetByID(ctx context.Context, id int64) (core.Note, error) {
	if id <= 0 {
		return core.Note{}, ErrInvalidNote
	}
	n, err := s.repo.GetByID(id)
	if err != nil {
		return core.Note{}, ErrNoteNotFound
	}
	return n, nil
}

func (s *noteService) List(ctx context.Context) ([]core.Note, error) {
	return s.repo.List()
}

func (s *noteService) Update(ctx context.Context, id int64, n core.Note) (core.Note, error) {
	if id <= 0 || strings.TrimSpace(n.Title) == "" {
		return core.Note{}, ErrInvalidNote
	}
	updated, err := s.repo.Update(id, n)
	if err != nil {
		return core.Note{}, ErrNoteNotFound
	}
	return updated, nil
}

func (s *noteService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return ErrInvalidNote
	}
	if err := s.repo.Delete(id); err != nil {
		return ErrNoteNotFound
	}
	return nil
}
