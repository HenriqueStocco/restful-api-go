package database

import (
	"database/sql"
)

func Migration(db *sql.DB) error {
	const schemas = /* sql */ `
		-- Create tables
		CREATE TABLE IF NOT EXISTS notes (
    		id TEXT PRIMARY KEY NOT NULL,
    		title TEXT NOT NULL,
    		description TEXT,
    		color TEXT DEFAULT "white",
			active INTEGER NOTE NULL DEFAULT 0,
    		created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		
		-- Create indexes
		CREATE INDEX IF NOT EXISTS idx_notes_id ON notes(id);
	`

	_, err := db.Exec(schemas)

	return err
}
