package notes

import "errors"

func FindAllNotesService() (*[]NoteSchema, error) {
	notesData, err := FindAllNotesRepository()

	if err != nil {
		return nil, errors.New("No notes found")
	}

	return notesData, nil
}
