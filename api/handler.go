package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/victorabarros/work-at-olist/internal/database"
)

// SetUpRoutes set api routes
func SetUpRoutes(r *mux.Router, db *database.Authors, myDb *database.Database) {
	r.HandleFunc("/authors", listAuthors(db)).Methods(http.MethodGet)
	r.HandleFunc("/authors/{id:[0-9]+}", getAuthor(db)).Methods(http.MethodGet)
	r.HandleFunc("/books", createBook(myDb)).Methods(http.MethodPost)
	// TODO add liveness and probeness
}
