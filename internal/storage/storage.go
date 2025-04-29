package storage

type Storage interface {
	CreateStudent(name string, email string, age uint8) (int64, error)
}
