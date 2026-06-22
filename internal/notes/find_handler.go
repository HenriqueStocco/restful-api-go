package notes

import (
	"net/http"
	"strings"

	h "github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
)

func FindAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.WriteJSON(w, http.StatusMethodNotAllowed, h.APIResponse{
			Success: false,
			Message: "Method not allowed",
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

func FindNoteByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.WriteJSON(w, http.StatusMethodNotAllowed, h.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use GET",
		})
		return
	}

	urlPath := strings.Split(r.URL.Path, "/")
	noteId := urlPath[len(urlPath)-1]

	noteData, err := FindNoteByIdService(noteId)

	if err != nil {
		h.WriteJSON(w, http.StatusNotFound, h.APIResponse{
			Success: false,
			Message: "Not note found",
		})
		return
	}

	h.WriteJSON(w, http.StatusOK, h.APIResponse{
		Success: true,
		Message: "Note data",
		Data:    noteData,
	})
}
