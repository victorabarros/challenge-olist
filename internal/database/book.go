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

// InsertBook inserts new book to db
func (db *Database) InsertBook(book Book) (int64, error) {
	query := fmt.Sprintf(
		`INSERT INTO books (name, edition, published_year)
		 VALUES ('%s', %d, %d);`,
		book.Name, book.Edition, book.PublicationYear,
	)

	resp := db.Connection.MustExec(query)

	return resp.LastInsertId()
}

// InsertBookAuthors inserts new book to db
func (db *Database) InsertBookAuthors(bookID int, authorIDs []int) error {
	values := fmt.Sprintf("(%d, %d)", bookID, authorIDs[0])

	for _, authorID := range authorIDs[1:] {
		values = fmt.Sprintf("%s , (%d, %d)", values, bookID, authorID)
	}

	query := fmt.Sprintf(
		`INSERT INTO books_authors (book_id, author_id) VALUES %s;`,
		values,
	)
	fmt.Println(query)

	resp := db.Connection.MustExec(query)
	fmt.Println(resp)

	return nil
}
