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

// filters are the conditions send by parameters on request
type filters map[string][]string

func createBook(b business.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"createBook\" route")
		var book database.Book

		err := json.NewDecoder(req.Body).Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{err.Error()})
			return
		}

		err = b.Create(book)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{"Fail connection on DB."})
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response{"Book created with success. ID: "}) // TODO add id
	}
}

// listBooks return list
func listBooks(b business.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// TODO rota não está funcnionando
		fmt.Println("Starting \"listBooks\" route")

		filters, err := extractFilters(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{err.Error()})
			return
		}

		books, err := b.List(filters)

		switch {
		case err != nil:
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{"Fail connection on DB."})
			return
		case books == nil: // TODO or len(books) == 0
			w.WriteHeader(http.StatusNotFound)
			return
		default:
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(books)
		}
	}
}

func extractFilters(req *http.Request) (f filters, err error) {
	// jsonschema validator https://github.com/xeipuuv/gojsonschema
	params := req.URL.Query()
	var e error

	publications, prs := params["publication"]
	if prs {
		for _, val := range publications {
			_, e = strconv.Atoi(val)
			if e != nil {
				err = fmt.Errorf("%s\r\n%s", err, e.Error())
			}
		}
		f["PublicationYears"] = publications
	}

	editions, prs := params["edition"]
	if prs {
		for _, val := range editions {
			_, e = strconv.Atoi(val)
			if e != nil {
				err = fmt.Errorf("%s\r\n%s", err, e.Error())
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
	return
}

func getBook(b business.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"getBook\" route")

		params := mux.Vars(req)
		id, _ := strconv.Atoi(params["id"])

		author, err := b.Get(id)

		switch {
		case err != nil:
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{"Fail connection on DB."})
			return
		case author == nil:
			w.WriteHeader(http.StatusNotFound)
			return
		default:
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(author)
		}
	}
}

// putBook return list
func putBook(b business.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"putBook\" route")

		var err error
		params := mux.Vars(req)
		bookID, _ := strconv.Atoi(params["id"])
		// TODO check if bookID exist! bug http://localhost:8092/books/1091

		book := database.Book{}
		err = json.NewDecoder(req.Body).Decode(&book)

		if book.Edition == 0 {
			// TODO should use err.Error()?
			err = fmt.Errorf("%s\r\n%s", err, "edition can't be null.")
		}
		if book.PublicationYear == 0 {
			err = fmt.Errorf("%s\r\n%s", err, "publication_year can't be null.")
		}
		if book.Authors == nil || len(book.Authors) == 0 {
			err = fmt.Errorf("%s\r\n%s", err, "authors can't be null or empty.")
		}

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{err.Error()})
			return
		}

		book.ID = bookID
		err = b.Update(book)
		if err != nil {
			// TODO log error
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{"Fail connection on DB."})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		// TODO improve message response
		json.NewEncoder(w).Encode(bookID)
	}
}

// patchBook return list
func patchBook(b business.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"patchBook\" route")

		var err error
		params := mux.Vars(req)
		bookID, _ := strconv.Atoi(params["id"])
		// TODO check if bookID exists, if not: status not found

		book := database.Book{}
		err = json.NewDecoder(req.Body).Decode(&book)
		if book.Authors != nil && len(book.Authors) == 0 {
			err = fmt.Errorf("%s\r\n%s", err, "authors can't be null or empty.")
		}

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{err.Error()})
			return
		}

		book.ID = bookID

		err = b.PartialUpdate(book)
		if err != nil {
			// TODO log error
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{"Fail connection on DB."})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		// TODO improve message response
		json.NewEncoder(w).Encode(bookID)
	}
}

// deleteBook return list
func deleteBook(b business.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"deleteBook\" route")

		params := mux.Vars(req)
		bookID, _ := strconv.Atoi(params["id"])
		// TODO É necessário validar se o bookID existe? status not found

		err := b.Delete(bookID)
		if err != nil {
			// TODO log error
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{"Fail connection on DB."})
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
