package database

import (
	"database/sql"
	"fmt"
)

func RunMigrations(db *sql.DB) error {
	const schemas = /* sql */ `
		-- Create tables
		CREATE TABLE IF NOT EXISTS notes (
    		id TEXT PRIMARY KEY NOT NULL,
    		title TEXT NOT NULL,
    		description TEXT NOT NULL,
    		color TEXT NOT NULL DEFAULT "white",
			active INTEGER NOTE NULL DEFAULT 1,
    		created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		
		-- Create indexes
		CREATE INDEX IF NOT EXISTS idx_notes_id ON notes(id);
	`

	if _, err := db.Exec(schemas); err != nil {
		return fmt.Errorf("failed to run database migrations: %w", err)
	}

	return nil
}
