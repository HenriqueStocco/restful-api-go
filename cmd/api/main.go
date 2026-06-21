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
	http.HandleFunc(s.SetURL("/health"), h.HealthHandler)
	http.HandleFunc(s.SetURL("/notes"), n.CreateNoteHandler)
	http.HandleFunc(s.SetURL("/notes/all"), n.FindAllNotesHandler)

	err := http.ListenAndServe("localhost:8000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
		return
	} else if err != nil {
		fmt.Printf("Error to starting the server: %v\n", err)

		os.Exit(1)
	}
}
