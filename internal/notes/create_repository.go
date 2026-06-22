package notes

import (
	"log"

	d "github.com/HenriqueStocco/restful-api-crud-go/internal/db"
	u "github.com/google/uuid"
)

func CreateNoteRepository(noteData *NoteSchema) (*NoteSchema, error) {
	db, err := d.DBConnection()

	if err != nil {
		log.Fatal("Failed to configure database connection: ", err)
	}

	defer db.Close()

	id := u.New()

	_, exErr := db.Exec(`
		INSERT INTO notes (
			id, 
			title, 
			description, 
			color
		) 
		VALUES (
			?, 
			?, 
			?, 
			?
		)`,
		id, noteData.Title, noteData.Description, noteData.Color,
	)

	if exErr != nil {
		return nil, exErr
	}

	createdNote, fErr := FindNoteByIdRepository(id)

	if fErr != nil {
		return nil, fErr
	}

	return createdNote, nil
}
