package api

import (
	"encoding/json"
	"fmt"
	"net/http"

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
