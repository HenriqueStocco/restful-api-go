package notes

import (
	"errors"

	"github.com/google/uuid"
)

func FindAllNotesService() (*[]NoteSchema, error) {
	notesData, err := FindAllNotesRepository()

	if err != nil {
		return nil, errors.New("No notes found")
	}

	return notesData, nil
}

func FindNoteByIdService(noteId string) (*NoteSchema, error) {
	noteIdUUID, parseError := uuid.Parse(noteId)

	if parseError != nil {
		return nil, errors.New("Error to parse noteId string to UUID")
	}

	noteData, err := FindNoteByIdRepository(noteIdUUID)

	if err != nil {
		return nil, errors.New("No note found")
	}

	return noteData, nil
}
