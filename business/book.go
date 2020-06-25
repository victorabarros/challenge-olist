package business

import (
	"fmt"

	"github.com/victorabarros/challenge-olist/internal/database"
)

// Book business layer
type Book struct {
	// TODO faz sentido usar uma interface? Ficaria bem grande
	DB database.Database
}

// Create inserts new book to db
func (b *Book) Create(book database.Book) error {
	for _, id := range book.Authors {
		// TODO mover esta validação para camada anterior, para tratar a resposta adequadamente
		a, err := b.DB.GetAuthorByID(id)
		if err != nil {
			// 503
			return err
		} else if a == nil {
			// 400
			return fmt.Errorf("Author id %d doesn't exist", id)
		}
	}

	id, err := b.DB.InsertBook(book)
	if err != nil {
		return err
	}

	book.ID = int(id)

	err = b.DB.InsertBookAuthors(book.ID, book.Authors)
	if err != nil {
		return err
	}
	return nil
}

// List books after filter
func (b *Book) List(filters map[string][]string) ([]database.Book, error) {
	ans, err := b.DB.ListBooks(filters)
	if err != nil {
		return nil, err
	}

	books := ans

	for idx, book := range books {
		authors, _ := b.DB.GetAuthorsIDByBookID(book.ID)
		books[idx].Authors = authors
	}
	return books, nil
}

// Get books after filter
func (b *Book) Get(id int) (*database.Book, error) {
	book, err := b.DB.GetBookByID(id)
	if err != nil {
		return nil, err
	}

	authors, _ := b.DB.GetAuthorsIDByBookID(book.ID)
	if err != nil {
		return nil, err
	}
	book.Authors = authors

	return book, nil
}

// Delete books after filter
func (b *Book) Delete(id int) error {
	return b.DB.DeleteBook(id)
}

// Update books after filter
func (b *Book) Update(book database.Book) error {
	err := b.DB.UpdateBook(book)
	if err != nil {
		return err
	}

	err = b.DB.DeleteBookAuthors(book.ID)
	if err != nil {
		return err
	}

	err = b.DB.InsertBookAuthors(book.ID, book.Authors)
	if err != nil {
		return err
	}

	return nil
}

// PartialUpdate books after filter
func (b *Book) PartialUpdate(book database.Book) error {
	err := b.DB.PartialUpdateBook(book)
	if err != nil {
		return err
	}

	if book.Authors != nil && len(book.Authors) != 0 {
		err = b.DB.DeleteBookAuthors(book.ID)
		if err != nil {
			return err
		}

		err = b.DB.InsertBookAuthors(book.ID, book.Authors)
		if err != nil {
			return err
		}
	}

	return nil
}
