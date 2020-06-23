package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/victorabarros/challenge-olist/business"
	"github.com/victorabarros/challenge-olist/internal/database"
)

// listAuthors return list
func listAuthors(a *business.Author) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"listAuthors\" route")

		offset, limit, name, err := validateListQueryParams(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
				fmt.Println(err)
			}
			return
		}

		authors, err := a.List(offset, limit, name)

		switch {
		case err != nil:
			w.WriteHeader(http.StatusServiceUnavailable)
			// TODO log err
			// TODO build error response message
		case len(*authors) == 0:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(authors); err != nil {
				fmt.Println(err)
			}
		}
	}
}

// TODO improve name
func validateListQueryParams(req *http.Request) (
	offset int, limit int, name string, err error) {
	// jsonschema validator https://github.com/xeipuuv/gojsonschema
	params := req.URL.Query()

	nameQueue, prs := params["name"]
	if prs {
		// TODO add handler if len > 1
		name = nameQueue[0]
	}

	var e error

	limitStr, prs := params["limit"]
	if !prs {
		// Default value
		limitStr = []string{"100"} // TODO move to env
	}
	if len(limitStr) > 1 {
		err = fmt.Errorf("%s\n%s", err, "Only one parameter is valid to 'limit'")
	} else {
		limit, e = strconv.Atoi(limitStr[0])
		if e != nil {
			err = fmt.Errorf("%s\n%s", err, e)
		} else if limit < 0 {
			err = fmt.Errorf("%s\n%s", err, "'limit' parameter must be positive")
			limit = *new(int)
		}
	}

	offsetStr, prs := params["offset"]
	if !prs {
		// Default value
		offsetStr = []string{"0"} // TODO env
	}
	if len(offsetStr) > 1 {
		err = fmt.Errorf("%s\n%s", err, "Only one parameter is valid to 'offset'")
	} else {
		offset, e = strconv.Atoi(offsetStr[0])
		if e != nil {
			err = fmt.Errorf("%s\n%s", err, e)
		} else if offset < 0 {
			err = fmt.Errorf("%s\n%s", err, "'offset' parameter must be positive")
			offset = *new(int)
		}
	}

	return
}

func getAuthorByID(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"getAuthor\" route")

		params := mux.Vars(req)
		id, _ := strconv.Atoi(params["id"])

		author, err := db.GetAuthorByID(id)
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
