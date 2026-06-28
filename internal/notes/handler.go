package notes

import (
	"encoding/json"
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
		helpers.WriteJSON(w, http.StatusMethodNotAllowed, helpers.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use POST",
		})
		return
	}

	defer r.Body.Close()

	var noteData CreateNoteRequestBodytDTO

	if err := json.NewDecoder(r.Body).Decode(&noteData); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.APIResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	serviceErr := h.service.CreateNote(CreateNoteRequestBodytDTO{
		Title:       noteData.Title,
		Description: noteData.Description,
		Color:       noteData.Color,
	})

	if serviceErr != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.APIResponse{
			Success: false,
			Message: "Failed to create note",
			Error:   serviceErr.Error(),
		})
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, helpers.APIResponse{
		Success: true,
		Message: "Note created successfully",
	})
}
