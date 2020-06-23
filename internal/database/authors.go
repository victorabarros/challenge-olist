package database

import (
	"fmt"
)

// Author model
type Author struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

// CreateAuthor create author
func (db *Database) CreateAuthor(name string) (bool, error) {
	query := `INSERT INTO authors (name) VALUES (:name);`

	// TODO look for a other method that return error
	_, err := db.Connection.NamedExec(query,
		map[string]interface{}{
			"name": name,
		},
	)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	// TODO Se já tiver no banco, ignorar erro.

	return true, nil
}

// ListAuthors return a sub section
func (db *Database) ListAuthors(limit, offset int) ([]Author, error) { // TODO improve name
	query := fmt.Sprintf(
		"SELECT id, name FROM authors LIMIT %d OFFSET %d;",
		limit,
		offset,
	)

	return db.selectAuthors(query)
}

// GetAuthorByName return simgle Author
func (db *Database) GetAuthorByName(name string) ([]Author, error) {
	// TODO Is necessary LIMIT 1?
	query := fmt.Sprintf(
		"SELECT id, name FROM authors WHERE name = '%s';",
		name,
	)
	// TODO: retornar apenas *Author
	return db.selectAuthors(query)
}

// GetAuthorByID return simgle Author
func (db *Database) GetAuthorByID(id int) ([]Author, error) {
	// TODO Is necessary LIMIT 1?
	query := fmt.Sprintf(
		"SELECT id, name FROM authors WHERE id = %d;",
		id,
	)
	// TODO: retornar apenas *Author
	return db.selectAuthors(query)
}

func (db *Database) selectAuthors(query string) ([]Author, error) {
	authors := []Author{}
	if err := db.Connection.Select(&authors, query); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return authors, nil
}
