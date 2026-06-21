package notes

import (
	"encoding/json"
	"net/http"

	u "github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
)

/* Função responsável por gerenciar requisições HTTP para a criação de notas */
func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		u.WriteJSON(w, http.StatusMethodNotAllowed, u.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use POST",
		})
		return
	}

	defer r.Body.Close()

	var noteData NoteSchema

	if err := json.NewDecoder(r.Body).Decode(&noteData); err != nil {
		u.WriteJSON(w, http.StatusBadRequest, u.APIResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	result, err := CreateNoteService(&noteData)

	if err != nil {
		u.WriteJSON(w, http.StatusNotFound, u.APIResponse{
			Success: false,
			Message: "Failed to create note",
			Error:   err.Error(),
		})
		return
	}

	u.WriteJSON(w, http.StatusCreated, u.APIResponse{
		Success: true,
		Message: "Note created successfully",
		Data:    result,
	})
}
