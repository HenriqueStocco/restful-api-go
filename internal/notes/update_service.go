package notes

import (
	"errors"
	"strings"
)

func UpdateNoteService(reqUrl string, noteData *NoteSchema) (bool, error) {
	if reqUrl == "" {
		return false, errors.New("Empty request url")
	}

	requestParam := strings.Split(reqUrl, "/")
	noteIdString := requestParam[len(requestParam)-1]
	// noteId, parseError := uuid.Parse(noteIdString)

	// if parseError != nil {
	// 	return false, errors.New("Error to parse noteId string to UUID")
	// }

	data := NoteSchema{
		ID:          noteIdString,
		Title:       noteData.Title,
		Description: noteData.Description,
		Color:       noteData.Color,
	}

	result, updateError := UpdateNoteRepository(&data)

	if updateError != nil {
		return false, errors.New("Failed to update note")
	}

	if result != true {
		return false, errors.New("Failed to update note")
	}

	return result, nil
}
