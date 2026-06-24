package notes

import (
	"errors"
	"strings"
)

func DeleteNoteByIdService(url string) (bool, error) {
	if url == "" {
		return false, errors.New("URL is empty")
	}

	pathParam := strings.Split(strings.Trim(url, " "), "/")
	noteId := pathParam[len(pathParam)-1]

	result, repositoryErr := DeleteNoteByIdRepository(noteId)

	if repositoryErr != nil || result != true {
		return false, errors.New(repositoryErr.Error())
	}

	return result, nil
}

func DeleteAllNotesService() (bool, error) {
	result, repositoryErr := DeleteAllNotesRepository()

	if repositoryErr != nil || result != true {
		return false, errors.New(repositoryErr.Error())
	}

	return result, nil
}
