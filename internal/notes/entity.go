package notes

import (
	"strings"

	"github.com/google/uuid"
)

type NoteEntity struct {
	ID          uuid.UUID
	Title       string
	Description string
	Color       string
	Active      bool
}

func (n *NoteEntity) Validate() error {
	n.Title = strings.TrimSpace(n.Title)
	n.Description = strings.TrimSpace(n.Description)
	n.Color = strings.TrimSpace(n.Color)

	if n.Title == "" {
		return ErrInvalidNoteTitle
	}
	if len(n.Title) < 4 {
		return ErrInvalidNoteTitle
	}

	if n.Description == "" {
		return ErrInvalidNoteDescription
	}
	if len(n.Description) < 8 {
		return ErrInvalidNoteDescription
	}

	if n.Color == "" {
		return ErrInvalidNoteColor
	}

	return nil
}

func NewNoteEntity(id uuid.UUID, title string, description string, color string) (*NoteEntity, error) {
	note := &NoteEntity{
		ID:          id,
		Title:       title,
		Description: description,
		Color:       color,
		Active:      true,
	}

	if err := note.Validate(); err != nil {
		return nil, err
	}

	return note, nil
}
