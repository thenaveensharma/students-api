package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thenaveensharma/students-api/internal/config"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	result, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS students 
	(id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	age INTEGER NOT NULL)`)

	if err != nil {
		return nil, err
	}

	fmt.Println("student table created successfuly", result)

	return &Sqlite{Db: db}, nil
}
