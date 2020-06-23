package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/victorabarros/challenge-olist/business"

	"github.com/gorilla/mux"
	"github.com/victorabarros/challenge-olist/internal/database"
)

type filters map[string][]string

func createBook(b business.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"createBook\" route")
		var book database.Book

		err := json.NewDecoder(req.Body).Decode(&book)
		if err != nil {
			// TODO make the same at /api/author
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = b.Create(book)
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

		filters, errors := extractFilters(req)
		if len(errors) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(struct{ messages []error }{errors})
			return
		}

		ans, err := db.ListBooks(filters)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			// TODO add message response
			return
		}
		books := ans

		for idx, book := range books {
			authors, _ := db.GetAuthorsIDByBookID(book.ID)
			books[idx].Authors = authors
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).
			Encode(books); err != nil {
			fmt.Println(err)
		}
	}
}

func extractFilters(req *http.Request) (filters, []error) {
	// jsonschema validator https://github.com/xeipuuv/gojsonschema
	params := req.URL.Query()
	f := make(filters)
	errors := []error{}

	publications, prs := params["publication"]
	if prs {
		for _, val := range publications {
			_, err := strconv.Atoi(val)
			if err != nil {
				errors = append(errors, err)
			}
		}
		f["PublicationYears"] = publications
	}

	editions, prs := params["edition"]
	if prs {
		for _, val := range editions {
			_, err := strconv.Atoi(val)
			if err != nil {
				errors = append(errors, err)
			}
		}
		f["Editions"] = editions
	}

	authorsIDs, prs := params["author"]
	if prs {
		f["Authors"] = authorsIDs
	}

	names, prs := params["name"]
	if prs {
		f["Names"] = names
	}

	return f, errors
}

func getBookByID(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"getBook\" route")

		params := mux.Vars(req)
		id, _ := strconv.Atoi(params["id"])

		author, err := db.GetBookByID(id)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		// TODO se len(author) == 0 http.StatusNotFound

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(author); err != nil {
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
