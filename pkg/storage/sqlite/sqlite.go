package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/SAHIL-Sharma21/students-management/pkg/config"
	"github.com/SAHIL-Sharma21/students-management/pkg/types"
	_ "github.com/mattn/go-sqlite3" //drivers for sqlite
)

// implemeting interface
type Sqlite struct {
	Db *sql.DB
}

// we will config in New
func New(cfg *config.Config) (*Sqlite, error) {
	//making db connection
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	//creating table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	age INT NOT NULL
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

// struct is implementing the interface
func (s *Sqlite) CreateStudent(name, email string, age int) (int64, error) {
	//inserting to db
	//query create
	query := `INSERT INTO students (name, email, age) VALUES (?, ?, ?);`

	stmt, err := s.Db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	//execute query
	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	//query
	query := `SELECT * FROM students WHERE id = ? LIMIT 1;`

	stmt, err := s.Db.Prepare(query)

	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	//serilizing with the student struct
	var student types.Student

	//execute query
	//scan will add the data from the database in struct
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with the given id: %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}

	return student, nil
}

// function to handle get list of students
func (s *Sqlite) GetListOfStudents() ([]types.Student, error) {

	query := `SELECT id, name, email, age FROM students;` //for production we have to use pagination

	stmt, err := s.Db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	//exec
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	//student var
	var students []types.Student

	for rows.Next() {
		var student types.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)

		if err != nil {
			return nil, err
		}

		//append in slice
		students = append(students, student)
	}

	return students, nil
}

func (s *Sqlite) UpdateStudent(id int64, name string, email string, age int) (int64, error) {
	query := `UPDATE students SET name = ?, email = ?, age = ? WHERE id = ? ;`

	stmt, err := s.Db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	//execute the query
	result, err := stmt.Exec(name, email, age, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("error updating studnent with id: %s", fmt.Sprint(id))
	}

	return rowsAffected, nil

}

func (s *Sqlite) DeleteStudentById(id int64) error {
	query := `DELETE FROM students WHERE id = ?;`

	stmt, err := s.Db.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	//execute the query
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	isDeleted, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if isDeleted == 0 {
		return fmt.Errorf("no student found with the given id: %s", fmt.Sprint(id))
	}

	return nil
}
