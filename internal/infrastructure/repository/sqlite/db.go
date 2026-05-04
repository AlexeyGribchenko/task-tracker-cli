package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	StoragePath string `env:"SQLITE_STORAGE_PATH" env-default:"./storage/task_storage.db"`
}

type Storage struct {
	db *sql.DB
}

func New(cfg Config) (*Storage, error) {

	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to sqlite3 db: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return fmt.Errorf("Failed to close db connection: db in nil!")
}
