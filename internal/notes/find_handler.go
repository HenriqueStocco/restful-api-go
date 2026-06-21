package notes

import (
	"net/http"

	h "github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
)

func FindAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.WriteJSON(w, http.StatusMethodNotAllowed, h.APIResponse{
			Success: false,
			Message: "Methodo not allowed",
			Error:   "use GET",
		})
		return
	}

	notesData, findErr := FindAllNotesService()

	if findErr != nil {
		h.WriteJSON(w, http.StatusNotFound, h.APIResponse{
			Success: false,
			Message: "No notes found",
			Error:   findErr.Error(),
		})
		return
	}

	h.WriteJSON(w, http.StatusOK, h.APIResponse{
		Success: true,
		Message: "Notes",
		Data:    notesData,
	})
}
