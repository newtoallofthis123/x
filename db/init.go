package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	db *sql.DB
}

// MakeDb initializes and returns a new database instance at the specified path.
func MakeDb(path string) (Db, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return Db{}, err
	}
	return Db{db}, nil
}

// Init initializes the database.
func (d *Db) Init() error {
	_, err := d.db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, name TEXT, cmd TEXT, last_used DATETIME)")
	return err
}

func (d *Db) Truncate() error {
	_, err := d.db.Exec("DELETE FROM tasks")
	return err
}
