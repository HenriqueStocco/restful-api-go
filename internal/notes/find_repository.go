package notes

import (
	d "github.com/HenriqueStocco/restful-api-crud-go/internal/db"
	u "github.com/google/uuid"
)

func FindAllNotesRepository() (*[]NoteSchema, error) {
	db, dbError := d.DBConnection()

	if dbError != nil {
		return nil, dbError
	}

	rows, queryError := db.Query(`
		SELECT 
			* 
		FROM notes;
	`)

	if queryError != nil {
		return nil, queryError
	}

	defer rows.Close()

	var notes []NoteSchema

	for rows.Next() {
		var note NoteSchema

		scanError := rows.Scan(
			&note.ID,
			&note.Title,
			&note.Description,
			&note.Color,
			&note.CreatedAt,
			&note.UpdatedAt,
		)

		if scanError != nil {
			return nil, scanError
		}

		notes = append(notes, note)
	}

	if rError := rows.Err(); rError != nil {
		return nil, rError
	}

	return &notes, nil
}

func FindNoteById(id u.UUID) (*NoteSchema, error) {
	var note NoteSchema

	db, dbError := d.DBConnection()

	if dbError != nil {
		return nil, dbError
	}

	queryError := db.QueryRow(`
		SELECT 
			id,
			title,
			description,
			color,
			created_at,
			updated_at 
		FROM notes
		WHERE id = ?;
	`, id).Scan(
		&note.ID,
		&note.Title,
		&note.Description,
		&note.Color,
		&note.CreatedAt,
		&note.UpdatedAt,
	)

	if queryError != nil {
		return nil, queryError
	}

	return &note, nil
}
