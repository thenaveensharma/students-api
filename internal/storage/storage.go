package storage

import (
	"github.com/thenaveensharma/students-api/internal/types"
)

type Storage interface {
	CreateStudent(name string, email string, age uint8) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetAllStudents() ([]types.Student, error)
	DeleteStudentById(id int64) error
}
