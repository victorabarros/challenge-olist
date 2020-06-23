package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/victorabarros/challenge-olist/internal/database"
)

func createBook(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"createBook\" route")
		var book database.Book

		err := json.NewDecoder(req.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, id := range book.Authors {
			// db.GetAuthorByID(id)
			// TODO chegar se authores existe no banco
			fmt.Println(id)
		}

		// TODO Validar se o livro já não existe
		id, err := db.InsertBook(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			// TODO mensagem de resposta
			return
		}
		book.ID = int(id)

		err = db.InsertBookAuthors(book.ID, book.Authors)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			// TODO mensagem de resposta
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		// TODO mensagem de resposta
	}
}

// listBooks return list
func listBooks(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"listBooks\" route")

		ans, err := db.ListBooks()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			// TODO add message response
			return
		}
		books := *ans

		for idx, book := range books {
			authors, _ := db.GetAuthorsIDByBookID(book.ID)
			books[idx].Authors = *authors
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).
			Encode(books); err != nil {
			fmt.Println(err)
		}
	}
}

// putBook return list
func putBook(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"putBook\" route")

		params := mux.Vars(req)
		bookID, _ := strconv.Atoi(params["id"])
		// TODO handler error
		// TODO check if bookID

		book := database.Book{}
		err := json.NewDecoder(req.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// TODO validar se edition, name, publication_year e len(authors) != 0
		// => badrequest
		book.ID = bookID

		_ = db.UpdateBook(book)
		// if err != nil {
		// 	fmt.Println(err)
		// 	http.Error(w, err.Error(), http.StatusServiceUnavailable)
		// 	// TODO mensagem de resposta
		// 	return
		// }

		_ = db.DeleteBookAuthors(book.ID)
		// if err != nil {
		// 	fmt.Println(err)
		// 	http.Error(w, err.Error(), http.StatusServiceUnavailable)
		// 	// TODO mensagem de resposta
		// 	return
		// }

		_ = db.InsertBookAuthors(book.ID, book.Authors)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusServiceUnavailable)
		// 	// TODO mensagem de resposta
		// 	return
		// }

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		// TODO improve message response
		if err := json.NewEncoder(w).Encode(bookID); err != nil {
			fmt.Println(err)
		}
	}
}

// patchBook return list
func patchBook(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"patchBook\" route")

		params := mux.Vars(req)
		bookID, _ := strconv.Atoi(params["id"])
		// TODO handler error
		// TODO check if bookID exists, if not: status not found

		book := database.Book{}
		err := json.NewDecoder(req.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		book.ID = bookID

		_ = db.PartialUpdateBook(book)
		// if err != nil {
		// 	fmt.Println(err)
		// 	http.Error(w, err.Error(), http.StatusServiceUnavailable)
		// 	// TODO mensagem de resposta
		// 	return
		// }

		if book.Authors != nil && len(book.Authors) != 0 {
			// TODO improve this update.
			// Is possible make with only one connection?
			_ = db.DeleteBookAuthors(book.ID)

			_ = db.InsertBookAuthors(book.ID, book.Authors)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		// TODO improve message response
		if err := json.NewEncoder(w).Encode(bookID); err != nil {
			fmt.Println(err)
		}
	}
}

// deleteBook return list
func deleteBook(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"deleteBook\" route")

		params := mux.Vars(req)
		bookID, _ := strconv.Atoi(params["id"])
		// TODO handler error
		// TODO check if bookID exists, if not: status not found

		_ = db.DeleteBook(bookID)

		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
	}
}
