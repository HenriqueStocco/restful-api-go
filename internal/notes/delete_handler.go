package notes

import (
	"net/http"

	h "github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
	u "github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
)

func DeleteNoteByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		u.WriteJSON(w, http.StatusMethodNotAllowed, u.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use DELETE",
		})
		return
	}

	_, serviceErr := DeleteNoteByIdService(r.URL.Path)

	if serviceErr != nil {
		h.WriteJSON(w, http.StatusInternalServerError, h.APIResponse{
			Success: false,
			Message: "Failed to delete note",
			Error:   serviceErr.Error(),
		})
		return
	}

	h.WriteJSON(w, http.StatusCreated, h.APIResponse{
		Success: true,
		Message: "Notes was deleted successfully",
	})
}

func DeleteAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		u.WriteJSON(w, http.StatusMethodNotAllowed, u.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use DELETE",
		})
		return
	}

	_, serviceErr := DeleteAllNotesService()

	if serviceErr != nil {
		h.WriteJSON(w, http.StatusInternalServerError, h.APIResponse{
			Success: false,
			Message: "Failed to delete all notes",
			Error:   serviceErr.Error(),
		})
		return
	}

	h.WriteJSON(w, http.StatusCreated, h.APIResponse{
		Success: true,
		Message: "All notes was deleted successfully",
	})
}
