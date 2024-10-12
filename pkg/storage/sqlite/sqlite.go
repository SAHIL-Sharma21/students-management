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
