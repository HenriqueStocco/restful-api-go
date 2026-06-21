package health

import (
	"net/http"

	u "github.com/HenriqueStocco/restful-api-crud-go/internal/helpers"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		u.WriteJSON(w, http.StatusMethodNotAllowed, u.APIResponse{
			Success: false,
			Message: "Method not allowed",
			Error:   "use GET",
		})
	}

	u.WriteJSON(w, http.StatusOK, u.APIResponse{
		Success: true,
		Message: "API Running!",
	})
}
