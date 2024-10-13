package storage

import "github.com/SAHIL-Sharma21/students-management/pkg/types"

// making interface - imortant
type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetListOfStudents() ([]types.Student, error)
}
