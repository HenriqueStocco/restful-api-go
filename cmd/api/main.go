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
	http.HandleFunc(string("GET "+s.SetURL("/health")), h.HealthHandler)
	http.HandleFunc(string("POST "+s.SetURL("/note/create")), n.CreateNoteHandler)
	http.HandleFunc(string("GET "+s.SetURL("/note/all")), n.FindAllNotesHandler)
	http.HandleFunc(string("GET "+s.SetURL("/note/{noteId}")), n.FindNoteByIdHandler)
	http.HandleFunc(string("PUT "+s.SetURL("/note/update/{noteId}")), n.UpdateNoteHandler)
	http.HandleFunc(string("PATCH "+s.SetURL("/note/update/title/{noteId}")), n.UpdateTitleNoteHandler)
	http.HandleFunc(string("PATCH "+s.SetURL("/note/update/description/{noteId}")), n.UpdateDescriptionNoteHandler)
	http.HandleFunc(string("PATCH "+s.SetURL("/note/update/color/{noteId}")), n.UpdateColorNoteHandler)
	http.HandleFunc(string("PATCH "+s.SetURL("/note/delete/{noteId}")), n.DeleteNoteByIdHandler)
	http.HandleFunc(string("PATCH "+s.SetURL("/note/delete")), n.DeleteAllNotesHandler)

	err := http.ListenAndServe("localhost:8000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
		return
	} else if err != nil {
		fmt.Printf("Error to starting the server: %v\n", err)

		os.Exit(1)
	}
}
