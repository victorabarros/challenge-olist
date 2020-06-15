package database

import "fmt"

// Book model
type Book struct {
	ID              int
	Name            string `json:"name"`
	Edition         int    `json:"edition"`
	PublicationYear int    `json:"publicationYear"`
	Authors         []int  `json:"authors"`
}

// CreateBook inserts new book to db
func (db *Database) CreateBook(book Book) (bool, error) {
	// Check if authors exists
	// if no return error
	// if yes insert book to `books`
	query := fmt.Sprintf(
		`INSERT INTO books (name, edition, published_year)
		VALUES (%s, %d, %d);`,
		book.Name, book.Edition, book.PublicationYear,
	)

	result := db.Connection.MustExec(query)
	fmt.Println(result)

	// retrive id and add relation on `books_authors`
	return true, nil
}
