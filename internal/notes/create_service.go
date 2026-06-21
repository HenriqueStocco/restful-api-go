package notes

import (
	"errors"
	"strings"
)

/* Responsável por tratar dados antes de criar uma nota */
func CreateNoteService(noteData *NoteSchema) (*NoteSchema, error) {
	noteData.Title = strings.TrimSpace(noteData.Title)
	noteData.Description = strings.TrimSpace(noteData.Description)
	noteData.Color = strings.TrimSpace(noteData.Color)

	if len(noteData.Title) < 4 {
		return nil, errors.New("Title must have at least 4 characters")
	}
	if len(noteData.Description) < 4 {
		return nil, errors.New("Description must have at least 4 characters")
	}

	if noteData.Color == "" {
		noteData.Color = "white"
	}

	result, err := CreateNoteRepository(noteData)

	if err != nil {
		return nil, err
	}

	return result, nil
}
