package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	d "github.com/HenriqueStocco/restful-api-crud-go/internal/database"
	s "github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
	"github.com/HenriqueStocco/restful-api-crud-go/internal/notes"
)

func main() {
	db, dbErr := d.DBConnection()

	if dbErr != nil {
		log.Fatal("failed to open database", dbErr)
	}

	repository := notes.NewNoteRepository(db)
	service := notes.NewNoteService(repository)
	handler := notes.NewNoteHandler(service)

	mux := http.NewServeMux()

	/* Routes */
	mux.HandleFunc(s.SetURL("POST", "/note"), handler.CreateNote)

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
