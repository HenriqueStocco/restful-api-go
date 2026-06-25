package notes

import (
	"errors"
	"strings"

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

func UpdateNoteService(url string, noteData *NoteSchema) (bool, error) {
	if url == "" {
		return false, errors.New("Empty request url")
	}

	requestParam := strings.Split(url, "/")
	noteIdString := requestParam[len(requestParam)-1]
	noteId, parseError := uuid.Parse(noteIdString)

	if parseError != nil {
		return false, errors.New("Error to parse noteId string to UUID")
	}

	data := NoteSchema{
		Id:          noteId,
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
