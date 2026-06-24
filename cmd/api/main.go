package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	h "github.com/HenriqueStocco/restful-api-crud-go/internal/health"
	s "github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
	n "github.com/HenriqueStocco/restful-api-crud-go/internal/notes"
)

func main() {
	mux := http.NewServeMux()

	/* GET routes */
	mux.HandleFunc(s.SetURL("GET", "/healthcheck"), h.HealthHandler)
	mux.HandleFunc(s.SetURL("GET", "/note"), n.FindAllNotesHandler)
	mux.HandleFunc(s.SetURL("GET", "/note/{noteId}"), n.FindNoteByIdHandler)
	/* POST routes */
	mux.HandleFunc(s.SetURL("POST", "/note"), n.CreateNoteHandler)
	/* PUT routes */
	mux.HandleFunc(s.SetURL("PUT", "/note/{noteId}"), n.UpdateNoteHandler)
	/* PATCH routes */
	mux.HandleFunc(s.SetURL("PATCH", "/note/{noteId}/title"), n.UpdateTitleNoteHandler)
	mux.HandleFunc(s.SetURL("PATCH", "/note/{noteId}/description"), n.UpdateDescriptionNoteHandler)
	mux.HandleFunc(s.SetURL("PATCH", "/note/{noteId}/color"), n.UpdateColorNoteHandler)
	/* DELETE routes */
	mux.HandleFunc(s.SetURL("DELETE", "/note/{noteId}"), n.DeleteNoteByIdHandler)
	mux.HandleFunc(s.SetURL("DELETE", "/note"), n.DeleteAllNotesHandler)

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
