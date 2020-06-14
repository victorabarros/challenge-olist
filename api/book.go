package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/victorabarros/work-at-olist/internal/database"
)

func createBook(db *database.Authors) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var book database.Book

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(req.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(book)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		// if err := json.NewEncoder(w).
		// 	Encode(db); err != nil {
		// 	fmt.Println(err)
		// }
	}
}
