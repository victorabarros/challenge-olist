package database

import "fmt"

// ListAuthorsByBookID return Authors by book ID
func (db *Database) ListAuthorsByBookID(id int) ([]Author, error) {
	query := fmt.Sprintf(
		`SELECT ba.author_id as id, a.name as name
		FROM books_authors ba LEFT JOIN authors a
		ON ba.author_id = a.id
		WHERE ba.book_id = %d;`,
		id,
	)
	// TODO: retornar apenas *Author
	return db.selectAuthors(query)
}

// GetAuthorsIDByBookID return a sub section
func (db *Database) GetAuthorsIDByBookID(bookID int) ([]int, error) {
	authors, err := db.ListAuthorsByBookID(bookID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	authorsID := make([]int, len(authors))
	for idx, author := range authors {
		authorsID[idx] = author.ID
	}

	return authorsID, nil
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

	// TODO look for a method that returns error
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
