package notes

import d "github.com/HenriqueStocco/restful-api-crud-go/internal/db"

func DeleteNoteByIdRepository(noteId string) (bool, error) {
	db, dbErr := d.DBConnection()

	if dbErr != nil {
		return false, dbErr
	}

	defer db.Close()

	const query string = `
        DELETE 
        FROM
            notes
        WHERE 
            id = ?;
    `

	_, queryErr := db.Exec(query, noteId)

	if queryErr != nil {
		return false, queryErr
	}

	return true, nil
}

func DeleteAllNotesRepository() (bool, error) {
	db, dbErr := d.DBConnection()

	if dbErr != nil {
		return false, dbErr
	}

	defer db.Close()

	const query string = `
        DELETE 
        FROM
            notes;
    `

	_, queryErr := db.Exec(query)

	if queryErr != nil {
		return false, queryErr
	}

	return true, nil
}
