package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/victorabarros/work-at-olist/internal/database"
)

// SetUpRoutes set api routes
func SetUpRoutes(r *mux.Router, db *database.Authors) {
	// TODO move to not especific file
	r.HandleFunc("/authors", listAuthors(db)).Methods(http.MethodGet)
	r.HandleFunc("/authors/{id:[0-9]+}", getAuthor(db)).Methods(http.MethodGet)
	r.HandleFunc("/books", createBook(db)).Methods(http.MethodPost)
	// TODO add liveness and probeness
}

// listAuthors return with offset (default = 0) and limit (default = 10) query params
func listAuthors(db *database.Authors) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		offset, limit, err := validateListQueryParams(req)
		// TODO;: Add name query param
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(struct{ Message error }{err})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).
			Encode(db.SubSection(offset, limit)); err != nil {
			fmt.Println(err)
		}
	}
}

func validateListQueryParams(req *http.Request) (int, int, error) { // TODO improve name
	// jsonschema validator https://github.com/xeipuuv/gojsonschema
	params := req.URL.Query()
	limit, prs := params["limit"]
	if !prs {
		limit = []string{"100"} // TODO move to env
	}

	offset, prs := params["offset"]
	if !prs {
		offset = []string{"0"}
	}

	limitHandled, err := strconv.Atoi(limit[0])
	if err != nil {
		return 0, 0, err
	}

	offsetHandled, err := strconv.Atoi(offset[0])
	if err != nil {
		return 0, 0, err
	}

	return offsetHandled, limitHandled, nil
}

func getAuthor(db *database.Authors) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		id, _ := strconv.Atoi(params["id"])

		val, prs := (*db)[id]
		if !prs {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		// TODO Isn't work:
		// >>> requests.get("http://localhost:8092/authors").headers
		// {'Date': 'Sun, 14 Jun 2020 00:07:29 GMT', 'Content-Length': '924', 'Content-Type': 'text/plain; charset=utf-8'}
		if err := json.NewEncoder(w).Encode(database.Authors{id: val}); err != nil {
			fmt.Println(err)
		}
	}
}
