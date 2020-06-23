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
		// TODO chegar se authores existe no banco
		// db.GetAuthorByID(id)
		fmt.Println(id)
	}

	// TODO Validar se o livro já não existe
	id, err := b.DB.InsertBook(book)
	if err != nil {
		return err
	}

	book.ID = int(id)
	// TODO how make both insertions at same time
	err = b.DB.InsertBookAuthors(book.ID, book.Authors)
	if err != nil {
		// handler if authorid doesn't exist, neste caso retornar 400 to client
		return err
	}
	return nil
}

// List books after filter
func (b *Book) List(filters map[string][]string) ([]database.Book, error) {
	// TODO investigar bug http://localhost:8092/books?name=gopl&author= resposta 503
	ans, err := b.DB.ListBooks(filters)
	if err != nil {
		return nil, err
	}

	books := ans

	for idx, book := range books {
		// TODO fazer essa busca usando business.Author
		authors, _ := b.DB.GetAuthorsIDByBookID(book.ID)
		books[idx].Authors = authors
	}
	return books, nil
}

// Get books after filter
func (b *Book) Get(id int) (*database.Book, error) {
	return b.DB.GetBookByID(id)
	// TODO adicionar authores à resposta
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
	// TODO é possível executar tudo numa transação só?
	// Minimizando a chance de erro

	return nil
}
