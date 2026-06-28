package notes

type CreateNoteRequestBodytDTO struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Color       string `json:"color,omitempty"`
}
