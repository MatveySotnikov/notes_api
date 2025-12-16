package repo

import "github.com/MatveySotnikov/notes-api/internal/core"

type NoteRepository interface {
	Create(note core.Note) (core.Note, error)
	GetByID(id int64) (core.Note, error)
	List() ([]core.Note, error)
	Update(id int64, note core.Note) (core.Note, error)
	Delete(id int64) error
}
