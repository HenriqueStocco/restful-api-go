package db

import (
	"database/sql"
	"log"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/* Gerenciador de banco de dados */
func DBConnection() (*sql.DB, error) {
	dbFilePath := filepath.Base("data/database.db")

	db, err := sql.Open("sqlite3", dbFilePath)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	if err := Migration(db); err != nil {
		log.Fatal("Failed to make migrations", err)
	}

	return db, nil
}
