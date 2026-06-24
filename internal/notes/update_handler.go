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

	newTitle := r.URL.Query().Get("value")

	if newTitle == "" {
		h.WriteJSON(w, http.StatusBadRequest, h.APIResponse{
			Success: false,
			Message: "Missing new title query param",
			Error:   "pass the query ?value=newtitle",
		})
		return
	}

	_, serviceErr := UpdateNoteTitleService(r.URL.Path, newTitle)

	if serviceErr != nil {
		h.WriteJSON(w, http.StatusInternalServerError, h.APIResponse{
			Success: false,
			Message: "Failed to update note title",
			Error:   serviceErr.Error(),
		})
		return
	}

	h.WriteJSON(w, http.StatusOK, h.APIResponse{
		Success: true,
		Message: "Note title was updated successfully",
	})
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

	newDescription := r.URL.Query().Get("value")

	if newDescription == "" {
		h.WriteJSON(w, http.StatusBadRequest, h.APIResponse{
			Success: false,
			Message: "Missing new description query param",
			Error:   "pass the query ?value=newdescription",
		})
		return
	}

	_, serviceErr := UpdateNoteDescriptionService(r.URL.Path, newDescription)

	if serviceErr != nil {
		h.WriteJSON(w, http.StatusInternalServerError, h.APIResponse{
			Success: false,
			Message: "Failed to update note description",
			Error:   serviceErr.Error(),
		})
		return
	}

	h.WriteJSON(w, http.StatusOK, h.APIResponse{
		Success: true,
		Message: "Note description was updated successfully",
	})
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

	newColor := r.URL.Query().Get("value")

	if newColor == "" {
		h.WriteJSON(w, http.StatusBadRequest, h.APIResponse{
			Success: false,
			Message: "Missing new color query param",
			Error:   "pass the query ?value=newcolor",
		})
		return
	}

	_, serviceErr := UpdateNoteColorService(r.URL.Path, newColor)

	if serviceErr != nil {
		h.WriteJSON(w, http.StatusInternalServerError, h.APIResponse{
			Success: false,
			Message: "Failed to update note color",
			Error:   serviceErr.Error(),
		})
		return
	}

	h.WriteJSON(w, http.StatusOK, h.APIResponse{
		Success: true,
		Message: "Note color was updated successfully",
	})
}
