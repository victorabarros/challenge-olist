package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/victorabarros/challenge-olist/business"
	"github.com/victorabarros/challenge-olist/internal/database"
)

// SetUpRoutes set api routes
func SetUpRoutes(r *mux.Router, a *business.Author, b business.Book, db *database.Database) {
	r.HandleFunc("/authors", listAuthors(a)).Methods(http.MethodGet)
	r.HandleFunc("/authors/{id:[0-9]+}", getAuthorByID(a)).Methods(http.MethodGet)
	r.HandleFunc("/books", createBook(b)).Methods(http.MethodPost)
	r.HandleFunc("/books", listBooks(db)).Methods(http.MethodGet)
	r.HandleFunc("/books/{id:[0-9]+}", getBookByID(db)).Methods(http.MethodGet)
	r.HandleFunc("/books/{id:[0-9]+}", putBook(db)).Methods(http.MethodPut)
	r.HandleFunc("/books/{id:[0-9]+}", patchBook(db)).Methods(http.MethodPatch)
	r.HandleFunc("/books/{id:[0-9]+}", deleteBook(db)).Methods(http.MethodDelete)
	// TODO add liveness and probeness
}
