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

		book := database.Book{}
		err := json.NewDecoder(req.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// TODO validar se edition, name, publication_year e len(authors) != 0

		params := mux.Vars(req)
		bookID, _ := strconv.Atoi(params["id"])
		// TODO handler error
		// TODO check if bookID , if not: status response 409 Conflict
		book.ID = bookID

		_ = db.UpdateBook(book, []string{"name", "edition", "publication_year"})
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

		if err := json.NewEncoder(w).
			Encode(bookID); err != nil {
			fmt.Println(err)
		}
	}
}

// patchBook return list
func patchBook(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"patchBook\" route")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}
}
