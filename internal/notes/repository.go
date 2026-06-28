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
