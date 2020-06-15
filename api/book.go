package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/victorabarros/work-at-olist/internal/database"
)

func createBook(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var book database.Book

		err := json.NewDecoder(req.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(book)

		_, err = db.CreateBook(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		// if err := json.NewEncoder(w).
		// 	Encode(db); err != nil {
		// 	fmt.Println(err)
		// }
	}
}
