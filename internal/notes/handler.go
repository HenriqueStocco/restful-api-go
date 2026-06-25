package notes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
)

/* Função responsável para criação de notas */
func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.WriteJSON(w, http.StatusMethodNotAllowed, helpers.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use POST",
		})
		return
	}

	defer r.Body.Close()

	var noteData NoteSchema

	if err := json.NewDecoder(r.Body).Decode(&noteData); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.APIResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	result, err := CreateNoteService(&noteData)

	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, helpers.APIResponse{
			Success: false,
			Message: "Failed to create note",
			Error:   err.Error(),
		})
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, helpers.APIResponse{
		Success: true,
		Message: "Note created successfully",
		Data:    result,
	})
}

func FindAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.WriteJSON(w, http.StatusMethodNotAllowed, helpers.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use GET",
		})
		return
	}

	notesData, findErr := FindAllNotesService()

	if findErr != nil {
		helpers.WriteJSON(w, http.StatusNotFound, helpers.APIResponse{
			Success: false,
			Message: "No notes found",
			Error:   findErr.Error(),
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.APIResponse{
		Success: true,
		Message: "Notes",
		Data:    notesData,
	})
}

func FindNoteByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.WriteJSON(w, http.StatusMethodNotAllowed, helpers.APIResponse{
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
		helpers.WriteJSON(w, http.StatusNotFound, helpers.APIResponse{
			Success: false,
			Message: "Not note found",
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.APIResponse{
		Success: true,
		Message: "Note data",
		Data:    noteData,
	})
}

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helpers.WriteJSON(w, http.StatusMethodNotAllowed, helpers.APIResponse{
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
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.APIResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   decordeErr.Error(),
		})
		return
	}

	_, serviceError := UpdateNoteService(r.URL.Path, &noteData)

	if serviceError != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.APIResponse{
			Success: false,
			Message: "Failed to update note",
			Error:   serviceError.Error(),
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.APIResponse{
		Success: true,
		Message: "Note was updated successfully",
	})
}

func UpdateTitleNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		helpers.WriteJSON(w, http.StatusMethodNotAllowed, helpers.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use PATCH",
		})
		return
	}

	newTitle := r.URL.Query().Get("value")

	if newTitle == "" {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.APIResponse{
			Success: false,
			Message: "Missing new title query param",
			Error:   "pass the query ?value=newtitle",
		})
		return
	}

	_, serviceErr := UpdateNoteTitleService(r.URL.Path, newTitle)

	if serviceErr != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.APIResponse{
			Success: false,
			Message: "Failed to update note title",
			Error:   serviceErr.Error(),
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.APIResponse{
		Success: true,
		Message: "Note title was updated successfully",
	})
}

func UpdateDescriptionNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		helpers.WriteJSON(w, http.StatusMethodNotAllowed, helpers.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use PATCH",
		})
		return
	}

	newDescription := r.URL.Query().Get("value")

	if newDescription == "" {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.APIResponse{
			Success: false,
			Message: "Missing new description query param",
			Error:   "pass the query ?value=newdescription",
		})
		return
	}

	_, serviceErr := UpdateNoteDescriptionService(r.URL.Path, newDescription)

	if serviceErr != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.APIResponse{
			Success: false,
			Message: "Failed to update note description",
			Error:   serviceErr.Error(),
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.APIResponse{
		Success: true,
		Message: "Note description was updated successfully",
	})
}

func UpdateColorNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		helpers.WriteJSON(w, http.StatusMethodNotAllowed, helpers.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use PATCH",
		})
		return
	}

	newColor := r.URL.Query().Get("value")

	if newColor == "" {
		helpers.WriteJSON(w, http.StatusBadRequest, helpers.APIResponse{
			Success: false,
			Message: "Missing new color query param",
			Error:   "pass the query ?value=newcolor",
		})
		return
	}

	_, serviceErr := UpdateNoteColorService(r.URL.Path, newColor)

	if serviceErr != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.APIResponse{
			Success: false,
			Message: "Failed to update note color",
			Error:   serviceErr.Error(),
		})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.APIResponse{
		Success: true,
		Message: "Note color was updated successfully",
	})
}

func DeleteNoteByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helpers.WriteJSON(w, http.StatusMethodNotAllowed, helpers.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use DELETE",
		})
		return
	}

	_, serviceErr := DeleteNoteByIdService(r.URL.Path)

	if serviceErr != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.APIResponse{
			Success: false,
			Message: "Failed to delete note",
			Error:   serviceErr.Error(),
		})
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, helpers.APIResponse{
		Success: true,
		Message: "Notes was deleted successfully",
	})
}

func DeleteAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helpers.WriteJSON(w, http.StatusMethodNotAllowed, helpers.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use DELETE",
		})
		return
	}

	_, serviceErr := DeleteAllNotesService()

	if serviceErr != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.APIResponse{
			Success: false,
			Message: "Failed to delete all notes",
			Error:   serviceErr.Error(),
		})
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, helpers.APIResponse{
		Success: true,
		Message: "All notes was deleted successfully",
	})
}
