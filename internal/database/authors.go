package database

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

	names := make([]string, len(csvLines)-1)
	for idx, line := range csvLines[1:] {
		names[idx] = line[0]
	}

	db.CreateAuthors(names)
}

// CreateAuthors create author
func (db *Database) CreateAuthors(names []string) (bool, error) {
	query := fmt.Sprintf(
		`INSERT INTO authors (name) VALUES ('%s');`,
		strings.Join(names, "'), ('"),
	)

	// TODO look for a other method that return error
	_ = db.Connection.MustExec(query)

	return true, nil
}

// ListAuthors return a sub section
func (db *Database) ListAuthors(limit, offset int) (*[]Author, error) { // TODO improve name
	query := fmt.Sprintf(
		"SELECT id, name FROM authors LIMIT %d OFFSET %d;",
		limit,
		offset,
	)

	return db.selectAuthors(query)
}

// GetAuthorByName return simgle Author
func (db *Database) GetAuthorByName(name string) (*[]Author, error) {
	// TODO Is necessary LIMIT 1?
	query := fmt.Sprintf(
		"SELECT id, name FROM authors WHERE name = '%s';",
		name,
	)
	// TODO: retornar apenas *Author
	return db.selectAuthors(query)
}

// GetAuthorByID return simgle Author
func (db *Database) GetAuthorByID(id int) (*[]Author, error) {
	// TODO Is necessary LIMIT 1?
	query := fmt.Sprintf(
		"SELECT id, name FROM authors WHERE id = %d;",
		id,
	)
	// TODO: retornar apenas *Author
	return db.selectAuthors(query)
}

func (db *Database) selectAuthors(query string) (*[]Author, error) {
	authors := []Author{}
	err := db.Connection.Select(&authors, query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &authors, err
}
