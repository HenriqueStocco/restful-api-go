package notes

type CreateNoteRequestBodytDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
}
