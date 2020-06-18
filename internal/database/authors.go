package database

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Author model
type Author struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

// LoadCsv to authors
func (db *Database) LoadCsv(name string) {
	csvFile, err := os.Open(name)
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, line := range csvLines[1:] {
		db.CreateAuthor(line[0])
		// (*a)[idx] = author{Name: line[0]}
	}
}

// CreateAuthor create author
func (db *Database) CreateAuthor(name string) (bool, error) {
	query := fmt.Sprintf(
		`INSERT INTO authors (name) VALUES ('%s');`, name,
	)

	_ = db.Connection.MustExec(query) // TODO look for a other method that return error

	return true, nil
}

// SubSection return a sub section
func (db *Database) SubSection(limit, offset int) ([]Author, error) { // TODO improve name
	query := fmt.Sprintf(
		"SELECT id, name FROM authors LIMIT %d OFFSET %d;",
		limit,
		offset,
	)

	authors := make([]Author, limit)
	err := db.Connection.Select(&authors, query)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

// GetByName return simgle Author
func (db *Database) GetByName(name string) (*Author, error) {
	query := fmt.Sprintf(
		"SELECT id, name FROM authors WHERE name = %s;",
		name,
	)

	author := Author{}
	err := db.Connection.Select(&author, query)
	if err != nil {
		return nil, err
	}

	return &author, nil
}

// GetByID return simgle Author
func (db *Database) GetByID(id int) (*Author, error) {
	query := fmt.Sprintf(
		"SELECT id, name FROM authors WHERE id = %d;",
		id,
	)

	author := Author{}
	err := db.Connection.Select(&author, query)
	if err != nil {
		return nil, err
	}

	return &author, nil
}
