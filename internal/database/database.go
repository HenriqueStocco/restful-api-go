package database

import (
	"database/sql"
	"fmt"

	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/* Gerenciador de banco de dados */
func DBConnection() (*sql.DB, error) {
	const dirName string = "data"
	const fileName string = "database.db"
	var dbFilePath string = filepath.Join(dirName, fileName)

	_, err := os.ReadDir(dirName)

	if err != nil {
		if errr := os.Mkdir("data", 0755); errr != nil {
			fmt.Printf("Error: %v\n", errr)
		}
	}

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
