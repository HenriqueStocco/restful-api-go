package notes

import (
	"github.com/google/uuid"
)

type NoteService struct {
	repository *NoteRepository
}

func NewNoteService(repository *NoteRepository) *NoteService {
	return &NoteService{
		repository: repository,
	}
}

/* Responsável por tratar dados antes de criar uma nota */
func (s *NoteService) CreateNote(note CreateNoteRequestBodytDTO) error {

	noteId := uuid.New()

	noteData, entityErr := NewNoteEntity(noteId, note.Title, note.Description, note.Color)

	if entityErr != nil {
		return entityErr
	}

	noteModel := NoteEntityToModel(noteData)

	if repositoryErr := s.repository.Create(&noteModel); repositoryErr != nil {
		return repositoryErr
	}

	return nil
}
