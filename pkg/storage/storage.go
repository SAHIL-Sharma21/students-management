package storage

// making interface - imortant
type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
}
