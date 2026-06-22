package notes

import (
	d "github.com/HenriqueStocco/restful-api-crud-go/internal/db"
	u "github.com/google/uuid"
)

func UpdateFullNoteByIdRepository(noteId u.UUID) (*NoteSchema, error) {
	db, conError := d.DBConnection()

	if conError != nil {
		return nil, conError
	}

	updatedNote, updateErr := db.Exec(`
        UPDATE notes
        SET = ?, ?, ?
        WHERE id = ?;
    `)

	defer db.Close()
}
