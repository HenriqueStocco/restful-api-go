package notes

import (
	d "github.com/HenriqueStocco/restful-api-crud-go/internal/db"
)

func UpdateNoteRepository(noteData *NoteSchema) (bool, error) {
	db, conError := d.DBConnection()

	if conError != nil {
		return false, conError
	}

	_, updateErr := db.Exec(`
        UPDATE 
			notes
        SET 
			title = ?, 
			description = ?, 
			color = ?
        WHERE id = ?
		RETURNING id;
    `, noteData.Title, noteData.Description, noteData.Color, noteData.ID)

	if updateErr != nil {
		return false, updateErr
	}

	defer db.Close()

	return true, nil
}
func UpdateNoteTitleRepository(noteId string, noteTitle string) (bool, error) {
	db, conError := d.DBConnection()

	if conError != nil {
		return false, conError
	}

	_, updateErr := db.Exec(`
        UPDATE 
			notes
        SET 
			title = ?
        WHERE id = ?;
    `, noteTitle, noteId)

	if updateErr != nil {
		return false, updateErr
	}

	defer db.Close()

	return true, nil
}

func UpdateNoteDescriptionRepository(noteId string, noteDescription string) (bool, error) {
	db, conError := d.DBConnection()

	if conError != nil {
		return false, conError
	}

	_, updateErr := db.Exec(`
        UPDATE 
			notes
        SET 
			description = ?
        WHERE id = ?;
    `, noteDescription, noteId)

	if updateErr != nil {
		return false, updateErr
	}

	defer db.Close()

	return true, nil
}
func UpdateNoteColorRepository(noteId string, noteColor string) (bool, error) {
	db, conError := d.DBConnection()

	if conError != nil {
		return false, conError
	}

	_, updateErr := db.Exec(`
        UPDATE 
			notes
        SET 
			color = ?
        WHERE id = ?;
    `, noteColor, noteId)

	if updateErr != nil {
		return false, updateErr
	}

	defer db.Close()

	return true, nil
}
