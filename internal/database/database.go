package database

import (
	"database/sql"
	"errors"
	"fmt"

	"os"
	"path/filepath"

	"github.com/HenriqueStocco/restful-api-crud-go/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

func ensureSQLiteDir(connectionString string) error {
	dir := filepath.Dir(connectionString)

	if dir == "." {
		return nil
	}

	if err := os.Mkdir(dir, 0755); err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil
		}

		return fmt.Errorf("failed to create SQLite directory: %w", err)
	}

	return nil
}

func OpenConnection(config config.DatabaseConfig) (*sql.DB, error) {
	if config.Driver == "sqlite3" {
		if err := ensureSQLiteDir(config.ConnectionString); err != nil {
			return nil, err
		}
	}

	db, dbErr := sql.Open(config.Driver, config.ConnectionString)

	if dbErr != nil {
		return nil, fmt.Errorf("failed to open database: %w", dbErr)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.MaxLifeTime)

	if pingErr := db.Ping(); pingErr != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", pingErr)
	}

	return db, nil
}
