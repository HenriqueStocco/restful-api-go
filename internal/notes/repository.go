package notes

import (
	"database/sql"
	"errors"
	"log"

	"github.com/HenriqueStocco/restful-api-crud-go/internal/database"
	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) GetAll() (*[]NoteSchema, error) {
	var notes []NoteSchema
	const query = `
		SELECT
			*
		FROM notes;
	`

	result, dbErr := r.db.Query(query)

	if errors.Is(dbErr, sql.ErrNoRows) {
		return nil, errors.New("Notes not found")
	}

	if dbErr != nil {
		return nil, dbErr
	}

	defer r.db.Close()

	for result.Next() {
		var note NoteSchema

		scanErr := result.Scan(
			&note.Id,
			&note.Title,
			&note.Description,
			&note.Color,
			&note.CreatedAt,
			&note.UpdatedAt,
		)

		if scanErr != nil {
			return nil, scanErr
		}

		notes = append(notes, note)
	}

	if resultErr := result.Err(); resultErr != nil {
		return nil, resultErr
	}

	return &notes, nil
}

func (r *Repository) GetById(id uuid.UUID) (*NoteSchema, error) {
	const query = `
		SELECT
			*
		FROM notes;
	`
	var note NoteSchema

	result, dbErr := r.db.Query(query)

	if errors.Is(dbErr, sql.ErrNoRows) {
		return nil, errors.New("Note not found")
	}
	if dbErr != nil {
		return nil, dbErr
	}

	defer r.db.Close()

	scanErr := result.Scan(&note.Id, &note.Title, &note.Description, &note.Color, &note.CreatedAt, &note.UpdatedAt)

	if scanErr != nil {
		return nil, scanErr
	}
	if result.Err() != nil {
		return nil, result.Err()
	}

	return &note, nil
}

func (r *Repository) Create(note *CreateNoteDTO) error {}

func (r *Repository) Update(id uuid.UUID) error {}

func (r *Repository) UpdateTitle(id uuid.UUID) error {}

func (r *Repository) UpdateDescription(id uuid.UUID) error {}

func (r *Repository) UpdateColor(id uuid.UUID) error {}

func (r *Repository) DeleteById(id uuid.UUID) error {}

func (r *Repository) DeleteAll() error {}

func CreateNoteRepository(noteData *NoteSchema) (*NoteSchema, error) {
	db, err := database.DBConnection()

	if err != nil {
		log.Fatal("Failed to configure database connection: ", err)
	}

	defer db.Close()

	id := uuid.New()

	_, exErr := db.Exec(`
		INSERT INTO notes (
			id, 
			title, 
			description, 
			color
		) 
		VALUES (
			?, 
			?, 
			?, 
			?
		)`,
		id, noteData.Title, noteData.Description, noteData.Color,
	)

	if exErr != nil {
		return nil, exErr
	}

	createdNote, fErr := FindNoteByIdRepository(id)

	if fErr != nil {
		return nil, fErr
	}

	return createdNote, nil
}

func UpdateNoteRepository(noteData *NoteSchema) (bool, error) {
	db, conError := database.DBConnection()

	if conError != nil {
		return false, conError
	}

	_, updateErr := db.Exec(`
        UPDATE 
			notes
        SET 
			title = ?, 
			description = ?, 
			color = ?
        WHERE id = ?
		RETURNING id;
    `, noteData.Title, noteData.Description, noteData.Color, noteData.Id)

	if updateErr != nil {
		return false, updateErr
	}

	defer db.Close()

	return true, nil
}
func UpdateNoteTitleRepository(noteId string, newTitle string) (bool, error) {
	db, conError := database.DBConnection()

	if conError != nil {
		return false, conError
	}

	_, updateErr := db.Exec(`
        UPDATE 
			notes
        SET 
			title = ?
        WHERE id = ?;
    `, newTitle, noteId)

	if updateErr != nil {
		return false, updateErr
	}

	defer db.Close()

	return true, nil
}

func UpdateNoteDescriptionRepository(noteId string, newDescription string) (bool, error) {
	db, conError := database.DBConnection()

	if conError != nil {
		return false, conError
	}

	_, updateErr := db.Exec(`
        UPDATE 
			notes
        SET 
			description = ?
        WHERE id = ?;
    `, newDescription, noteId)

	if updateErr != nil {
		return false, updateErr
	}

	defer db.Close()

	return true, nil
}

func UpdateNoteColorRepository(noteId string, newColor string) (bool, error) {
	db, conError := database.DBConnection()

	if conError != nil {
		return false, conError
	}

	_, updateErr := db.Exec(`
        UPDATE 
			notes
        SET 
			color = ?
        WHERE id = ?;
    `, newColor, noteId)

	if updateErr != nil {
		return false, updateErr
	}

	defer db.Close()

	return true, nil
}

func DeleteNoteByIdRepository(noteId string) (bool, error) {
	db, dbErr := database.DBConnection()

	if dbErr != nil {
		return false, dbErr
	}

	defer db.Close()

	const query string = `
        DELETE 
        FROM
            notes
        WHERE 
            id = ?;
    `

	_, queryErr := db.Exec(query, noteId)

	if queryErr != nil {
		return false, queryErr
	}

	return true, nil
}

func DeleteAllNotesRepository() (bool, error) {
	db, dbErr := database.DBConnection()

	if dbErr != nil {
		return false, dbErr
	}

	defer db.Close()

	const query string = `
        DELETE 
        FROM
            notes;
    `

	_, queryErr := db.Exec(query)

	if queryErr != nil {
		return false, queryErr
	}

	return true, nil
}
