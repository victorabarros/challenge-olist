package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/victorabarros/challenge-olist/internal/database"
)

// SetUpRoutes set api routes
func SetUpRoutes(r *mux.Router, db *database.Database) {
	r.HandleFunc("/authors", listAuthors(db)).Methods(http.MethodGet)
	r.HandleFunc("/authors/{id:[0-9]+}", getAuthorByID(db)).Methods(http.MethodGet)
	r.HandleFunc("/books", createBook(db)).Methods(http.MethodPost)
	r.HandleFunc("/books", listBooks(db)).Methods(http.MethodGet)
	r.HandleFunc("/books/{id:[0-9]+}", putBook(db)).Methods(http.MethodPut)
	r.HandleFunc("/books/{id:[0-9]+}", patchBook(db)).Methods(http.MethodPatch)
	// TODO add liveness and probeness
}
