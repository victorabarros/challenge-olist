package database

import (
	"fmt"
	"strings"
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
		fmt.Println(err) // TODO log
		return false, err
	}

	return true, nil
}

// ListAuthors return a sub section
func (db *Database) ListAuthors(limit, offset int, names []string) ([]Author, error) {
	var sets string
	if len(names) > 0 {
		sets = fmt.Sprintf("WHERE name IN ('%s')",
			strings.Join(names, `', '`),
		)
	}

	query := fmt.Sprintf(
		"SELECT id, name FROM authors %s LIMIT %d OFFSET %d;",
		sets,
		limit,
		offset,
	)

	return db.selectAuthors(query)
}

// GetAuthorByID return simgle Author
func (db *Database) GetAuthorByID(id int) (*Author, error) {
	query := fmt.Sprintf(
		"SELECT id, name FROM authors WHERE id = %d;",
		id,
	)
	authors, err := db.selectAuthors(query)
	if len(authors) == 0 {
		return nil, err
	}
	author := authors[0]
	return &author, err
}

func (db *Database) selectAuthors(query string) ([]Author, error) {
	authors := []Author{}
	if err := db.Connection.Select(&authors, query); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return authors, nil
}
