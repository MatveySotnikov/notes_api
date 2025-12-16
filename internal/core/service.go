package core

// NoteService определяет методы, которые предоставляет слой бизнес-логики.
// Обработчики будут зависеть только от этого интерфейса.
type NoteService interface {
	Create(note Note) (int64, error)
	// Add other methods like GetByID, List, Update...
}

// NoteRepository определяет методы, которые предоставляет слой хранения данных.
// Сервис будет зависеть только от этого интерфейса.
type NoteRepository interface {
	Create(note Note) (int64, error)
	// Add other storage-specific methods
}
