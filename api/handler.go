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
	// TODO add liveness and probeness
}
