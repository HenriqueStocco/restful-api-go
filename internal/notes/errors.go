package notes

import "errors"

var (
	ErrInvalidNoteTitle       = errors.New("invalid note title")
	ErrInvalidNoteDescription = errors.New("invalid note description")
	ErrInvalidNoteColor       = errors.New("invalid note color")
	ErrNoteNotExist           = errors.New("note not exist")
	ErrMissingNoteTitle       = errors.New("missing note title")
	ErrMissingNoteColor       = errors.New("missing note color")
)
