package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	h "github.com/HenriqueStocco/restful-api-crud-go/internal/health"
	s "github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
	"github.com/HenriqueStocco/restful-api-crud-go/internal/notes"
)

func main() {
	mux := http.NewServeMux()

	/* GET routes */
	mux.HandleFunc(s.SetURL("GET", "/healthcheck"), h.HealthHandler)
	mux.HandleFunc(s.SetURL("GET", "/note"), notes.FindAllNotesHandler)
	mux.HandleFunc(s.SetURL("GET", "/note/{noteId}"), notes.FindNoteByIdHandler)
	/* POST routes */
	mux.HandleFunc(s.SetURL("POST", "/note"), notes.CreateNoteHandler)
	/* PUT routes */
	mux.HandleFunc(s.SetURL("PUT", "/note/{noteId}"), notes.UpdateNoteHandler)
	/* PATCH routes */
	mux.HandleFunc(s.SetURL("PATCH", "/note/{noteId}/title"), notes.UpdateTitleNoteHandler)
	mux.HandleFunc(s.SetURL("PATCH", "/note/{noteId}/description"), notes.UpdateDescriptionNoteHandler)
	mux.HandleFunc(s.SetURL("PATCH", "/note/{noteId}/color"), notes.UpdateColorNoteHandler)
	/* DELETE routes */
	mux.HandleFunc(s.SetURL("DELETE", "/note/{noteId}"), notes.DeleteNoteByIdHandler)
	mux.HandleFunc(s.SetURL("DELETE", "/note"), notes.DeleteAllNotesHandler)

	fmt.Println("Server running at localhost:8000")

	httpErr := http.ListenAndServe("localhost:8000", mux)

	if errors.Is(httpErr, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")

		return
	}

	if httpErr != nil {
		fmt.Printf("Error to starting the server: %v\n", httpErr)

		os.Exit(1)
	}
}
