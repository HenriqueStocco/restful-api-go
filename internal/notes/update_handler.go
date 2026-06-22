package notes

import (
	"net/http"

	h "github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
)

func UpdateFullNoteByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.WriteJSON(w, http.StatusMethodNotAllowed, h.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use PUT",
		})
		return
	}

	noteData, serviceError := UpdateFullNoteByIdService(r.URL.Path)

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
		Data:    noteData,
	})
}

func UpdateTitleNoteByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		h.WriteJSON(w, http.StatusMethodNotAllowed, h.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use PATCH",
		})
		return
	}
}
func UpdateDescriptionNoteByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		h.WriteJSON(w, http.StatusMethodNotAllowed, h.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use PATCH",
		})
		return
	}
}
func UpdateColorNoteByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		h.WriteJSON(w, http.StatusMethodNotAllowed, h.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use PATCH",
		})
		return
	}
}
