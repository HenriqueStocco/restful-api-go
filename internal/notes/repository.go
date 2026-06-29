package notes

import (
	"database/sql"
)

type NoteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) *NoteRepository {
	return &NoteRepository{
		db: db,
	}
}

func (r *NoteRepository) Create(note *NoteModel) error {
	query := `
		INSERT INTO notes (
			id,
			title,
			description,
			color,
			active
		)
		VALUES (
			?,
			?,
			?,
			?,
			?
		);
	`

	_, execErr := r.db.Exec(query,
		note.ID,
		note.Title,
		note.Description,
		note.Color,
		note.Active,
	)

	if execErr != nil {
		return execErr
	}

	return nil
}

func (r *NoteRepository) GetAll() ([]*NoteModel, error) {
	query := `
		SELECT
			*
		FROM 
			notes;
	`

	rows, queryRowErr := r.db.Query(query)

	if queryRowErr != nil {
		return nil, queryRowErr
	}

	defer rows.Close()

	var noteArrayData []*NoteModel

	for rows.Next() {
		var note *NoteModel

		scanErr := rows.Scan(&note.ID, &note.Title, &note.Description, &note.Color, &note.Active, &note.CreatedAt, &note.UpdatedAt)

		if scanErr != nil {
			return nil, scanErr
		}

		noteArrayData = append(noteArrayData, note)
	}

	if rowErr := rows.Err(); rowErr != nil {
		return nil, rowErr
	}

	return noteArrayData, nil
}
