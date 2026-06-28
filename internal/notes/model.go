package notes

import (
	"time"

	"github.com/google/uuid"
)

type NoteModel struct {
	ID          uuid.UUID
	Title       string
	Description string
	Color       string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NoteEntityToModel(note *NoteEntity) NoteModel {
	return NoteModel{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,
		Color:       note.Color,
		Active:      note.Active,
	}
}
