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
func UpdateNoteTitleRepository(noteId string, newTitle string) (bool, error) {
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
    `, newTitle, noteId)

	if updateErr != nil {
		return false, updateErr
	}

	defer db.Close()

	return true, nil
}

func UpdateNoteDescriptionRepository(noteId string, newDescription string) (bool, error) {
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
    `, newDescription, noteId)

	if updateErr != nil {
		return false, updateErr
	}

	defer db.Close()

	return true, nil
}
func UpdateNoteColorRepository(noteId string, newColor string) (bool, error) {
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
    `, newColor, noteId)

	if updateErr != nil {
		return false, updateErr
	}

	defer db.Close()

	return true, nil
}
