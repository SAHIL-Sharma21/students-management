package sqlite

import (
	"database/sql"

	"github.com/SAHIL-Sharma21/students-management/pkg/config"
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
	return 0, nil
}
