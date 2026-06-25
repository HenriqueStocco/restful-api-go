package notes

type CreateNoteDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type UpdateNoteByIdBodyDTO struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Color       string `json:"color,omitempty"`
}
