package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HenriqueStocco/restful-api-crud-go/internal/config"
	"github.com/HenriqueStocco/restful-api-crud-go/internal/database"
	"github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
	"github.com/HenriqueStocco/restful-api-crud-go/internal/notes"
)

func main() {
	cfg := config.Load()

	db, dbErr := database.OpenConnection(cfg.Database)

	if dbErr != nil {
		log.Fatal("failed to open database", dbErr)
	}

	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatal("failed to run migrations: ", err)
	}

	repository := notes.NewNoteRepository(db)
	service := notes.NewNoteService(repository)
	handler := notes.NewNoteHandler(service)

	mux := http.NewServeMux()

	/* Routes */
	mux.HandleFunc(helpers.SetURL("POST", "/note"), handler.CreateNote)

	serverAddress := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	fmt.Printf("Server running at http://%s\n", serverAddress)

	httpErr := http.ListenAndServe(serverAddress, mux)

	if errors.Is(httpErr, http.ErrServerClosed) {
		fmt.Println("Server closed")
		return
	}

	if httpErr != nil {
		fmt.Printf("Error to starting the server: %v\n", httpErr)
		os.Exit(1)
	}
}
