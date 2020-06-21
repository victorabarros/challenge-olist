package database

import "fmt"

// Book model
type Book struct {
	ID              int
	Name            string `json:"name"             db:"name"`
	Edition         int    `json:"edition"          db:"edition"`
	PublicationYear int    `json:"publication_year" db:"publication_year"`
	Authors         []int  `json:"authors"          db:"authors"`
}

// InsertBook inserts new book to db
func (db *Database) InsertBook(book Book) (int64, error) {
	query := fmt.Sprintf(
		`INSERT INTO books (name, edition, publication_year)
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

	_ = db.Connection.MustExec(query)

	return nil
}

// ListBooks return a sub section
func (db *Database) ListBooks() (*[]Book, error) {
	query := `
		SELECT b.id AS id, b.name AS name, b.edition AS edition
		,b.publication_year AS publication_year
		#,GROUP_CONCAT(ba.author_id) AS authors
		FROM books b
		#LEFT JOIN books_authors ba ON b.id = ba.book_id
		#GROUP BY b.id
		;`

	books := []Book{}
	err := db.Connection.Select(&books, query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &books, err
}
