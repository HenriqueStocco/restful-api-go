package notes

import (
	"errors"
	"strings"
)

func UpdateNoteService(url string, noteData *NoteSchema) (bool, error) {
	if url == "" {
		return false, errors.New("Empty request url")
	}

	requestParam := strings.Split(url, "/")
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

func UpdateNoteTitleService(url string, noteTitle string) (bool, error) {
	if url == "" {
		return false, errors.New("Empty request url")
	}

	requestParam := strings.Split(url, "/")
	noteIdString := requestParam[len(requestParam)-1]

	result, updateError := UpdateNoteTitleRepository(noteIdString, noteTitle)

	if updateError != nil {
		return false, errors.New("Failed to update note title")
	}

	if result != true {
		return false, errors.New("Failed to update note title")
	}

	return result, nil
}

func UpdateNoteDescriptionService(url string, noteDescription string) (bool, error) {
	if url == "" {
		return false, errors.New("Empty request url")
	}

	requestParam := strings.Split(url, "/")
	noteIdString := requestParam[len(requestParam)-1]

	result, updateError := UpdateNoteDescriptionRepository(noteIdString, noteDescription)

	if updateError != nil {
		return false, errors.New("Failed to update note description")
	}

	if result != true {
		return false, errors.New("Failed to update note description")
	}

	return result, nil
}

func UpdateNoteColorService(url string, noteColor string) (bool, error) {
	if url == "" {
		return false, errors.New("Empty request url")
	}

	requestParam := strings.Split(url, "/")
	noteIdString := requestParam[len(requestParam)-1]

	result, updateError := UpdateNoteColorRepository(noteIdString, noteColor)

	if updateError != nil {
		return false, errors.New("Failed to update note color")
	}

	if result != true {
		return false, errors.New("Failed to update note color")
	}

	return result, nil
}
