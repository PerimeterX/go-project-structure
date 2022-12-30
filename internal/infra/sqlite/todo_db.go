package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"todoapp/pkg/todoapp"
)

const (
	createTableQuery = `CREATE TABLE IF NOT EXISTS todo (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"date" DATE,
		"value" TEXT
	);`

	insertQuery = `INSERT INTO todo(date, value) VALUES (?, ?);`

	selectQuery = `SELECT * FROM todo;`
)

type TodoDB struct {
	db *sql.DB
}

func NewTodoDB(file string) (*TodoDB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	statement, err := db.Prepare(createTableQuery)
	if err != nil {
		return nil, err
	}

	_, err = statement.Exec()
	if err != nil {
		return nil, err
	}

	return &TodoDB{db: db}, nil
}

func (t *TodoDB) Close() error {
	return t.db.Close()
}

func (t *TodoDB) Insert(task *todoapp.Task) error {
	statement, err := t.db.Prepare(insertQuery)
	if err != nil {
		return err
	}

	_, err = statement.Exec(task.Date, task.Value)
	return err
}

func (t *TodoDB) All() ([]*todoapp.Task, error) {
	cursor, err := t.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}

	defer cursor.Close()
	var result []*todoapp.Task
	for cursor.Next() {
		task := &todoapp.Task{}
		err = cursor.Scan(&task.ID, &task.Date, &task.Value)
		if err != nil {
			return nil, err
		}

		result = append(result, task)
	}
	return result, nil
}
