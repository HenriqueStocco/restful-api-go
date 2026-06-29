package notes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
)

type NoteHandler struct {
	service *NoteService
}

func NewNoteHandler(service *NoteService) *NoteHandler {
	return &NoteHandler{
		service: service,
	}
}

/* Função responsável para criação de notas */
func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.WriteError(w, http.StatusMethodNotAllowed, "INVALID_HTTP_METHOD", "invalid http method")
		return
	}

	defer r.Body.Close()

	var noteData CreateNoteRequestBodytDTO

	if err := json.NewDecoder(r.Body).Decode(&noteData); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, "INVALID_JSON", "Invalid request body")
		return
	}

	serviceErr := h.service.CreateNote(CreateNoteRequestBodytDTO{
		Title:       noteData.Title,
		Description: noteData.Description,
		Color:       noteData.Color,
	})

	if serviceErr != nil {
		switch {
		case errors.Is(serviceErr, ErrInvalidNoteTitle):
			helpers.WriteError(w, http.StatusBadRequest, "INVALID_NOTE_TITLE", "invalid note title")

		case errors.Is(serviceErr, ErrInvalidNoteDescription):
			helpers.WriteError(w, http.StatusBadRequest, "INVALID_NOTE_DESCRIPTION", "invalid note description")

		case errors.Is(serviceErr, ErrInvalidNoteTitle):
			helpers.WriteError(w, http.StatusBadRequest, "INVALID_NOTE_COLOR", "invalid note color")
		default:
			helpers.WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "internal server error")
		}

		return
	}

	helpers.WriteSuccess(w, http.StatusCreated, "note created successfully", nil)
}

func (h *NoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.WriteError(w, http.StatusMethodNotAllowed, "INVALID_HTTP_METHOD", "invalid http method")
		return
	}
}
