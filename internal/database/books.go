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
	if len(authorIDs) == 0 {
		return nil
	}

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

// DeleteBookAuthors inserts new book to db
func (db *Database) DeleteBookAuthors(bookID int) error {
	query := fmt.Sprintf(
		`DELETE FROM books_authors WHERE book_id = %d;`,
		bookID,
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
	if err := db.Connection.Select(&books, query); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &books, nil
}

// UpdateBook return a sub section
func (db *Database) UpdateBook(book Book) error {
	// TODO valid fields
	sets := fmt.Sprintf(
		`name = "%s", edition = %d, publication_year = %d`,
		book.Name, book.Edition, book.PublicationYear,
	)

	query := fmt.Sprintf(
		`UPDATE books SET %s WHERE id = %d;`,
		sets, book.ID,
	)
	fmt.Println(query)

	_ = db.Connection.MustExec(query)

	return nil
}

// PartialUpdateBook return a sub section
func (db *Database) PartialUpdateBook(book Book) error {
	var sets string

	if book.Name != "" {
		sets = fmt.Sprintf(`%s name = '%s',`, sets, book.Name)
	}
	if book.Edition != 0 {
		sets = fmt.Sprintf(`%s edition = %d,`, sets, book.Edition)
	}
	if book.PublicationYear != 0 {
		sets = fmt.Sprintf(`%s publication_year = %d,`,
			sets, book.PublicationYear)
	}

	if sets == "" {
		// TODO return error or nil?
		fmt.Println("Any field to update")
		return nil
	}

	// remove last comma/character
	sets = sets[:len(sets)-1]

	query := fmt.Sprintf(
		`UPDATE books SET%s WHERE id = %d;`,
		sets, book.ID,
	)

	_ = db.Connection.MustExec(query)

	return nil
}
