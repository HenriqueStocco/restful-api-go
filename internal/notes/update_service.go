package notes

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

func UpdateFullNoteByIdService(reqUrl string) (*NoteSchema, error) {
	if reqUrl == "" {
		return nil, errors.New("Empty request url")
	}

	requestParam := strings.Split(reqUrl, "/")
	noteIdString := requestParam[len(requestParam)-1]
	noteId, parseError := uuid.Parse(noteIdString)

	if parseError != nil {
		return nil, errors.New("Error to parse noteId string to UUID")
	}

	updatedNote, updateError := UpdateFullNoteByIdRepository(noteId)

	if updateError != nil {
		return nil, errors.New("Failed to update note")
	}

	return updatedNote
}
