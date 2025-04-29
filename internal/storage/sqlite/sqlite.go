package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thenaveensharma/students-api/internal/config"
	"github.com/thenaveensharma/students-api/internal/types"
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

func (s *Sqlite) CreateStudent(name string, email string, age uint8) (int64, error) {
	stmt, err := s.Db.Prepare(`INSERT INTO students (name, email, age) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare(`SELECT id, name, email, age FROM students WHERE id = ? LIMIT 1`)
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id :%s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}

	return student, nil
}
