package notes

import (
	"encoding/json"
	"net/http"

	h "github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
	u "github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
)

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.WriteJSON(w, http.StatusMethodNotAllowed, h.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use PUT",
		})
		return
	}

	defer r.Body.Close()

	var noteData NoteSchema

	decordeErr := json.NewDecoder(r.Body).Decode(&noteData)

	if decordeErr != nil {
		u.WriteJSON(w, http.StatusBadRequest, u.APIResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   decordeErr.Error(),
		})
		return
	}

	_, serviceError := UpdateNoteService(r.URL.Path, &noteData)

	if serviceError != nil {
		h.WriteJSON(w, http.StatusInternalServerError, h.APIResponse{
			Success: false,
			Message: "Failed to update note",
			Error:   serviceError.Error(),
		})
		return
	}

	h.WriteJSON(w, http.StatusOK, h.APIResponse{
		Success: true,
		Message: "Note was updated successfully",
	})
}

func UpdateTitleNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		h.WriteJSON(w, http.StatusMethodNotAllowed, h.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use PATCH",
		})
		return
	}
}

func UpdateDescriptionNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		h.WriteJSON(w, http.StatusMethodNotAllowed, h.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use PATCH",
		})
		return
	}
}

func UpdateColorNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		h.WriteJSON(w, http.StatusMethodNotAllowed, h.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use PATCH",
		})
		return
	}
}
