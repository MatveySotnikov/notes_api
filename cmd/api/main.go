package main

import (
	"log"
	"net/http"

	"github.com/MatveySotnikov/notes-api/internal/core/service"
	httpx "github.com/MatveySotnikov/notes-api/internal/http"
	"github.com/MatveySotnikov/notes-api/internal/http/handlers"
	"github.com/MatveySotnikov/notes-api/internal/repo"
)

func main() {
	noteRepo := repo.NewNoteRepoMem()
	noteService := service.NewNoteService(noteRepo)

	h := &handlers.Handler{
		Service: noteService,
	}

	r := httpx.NewRouter(h)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
