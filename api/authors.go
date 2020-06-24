package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/victorabarros/challenge-olist/business"
)

// listAuthors return list
func listAuthors(a *business.Author) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"listAuthors\" route")

		limit, offset, names, err := validateListParams(req)
		if err != nil {
			// TODO log err
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{err.Error()})
			return
		}

		authors, err := a.List(limit, offset, names)

		switch {
		case err != nil:
			// TODO log err
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response{"Fail connection on DB."})
		case len(*authors) == 0:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(authors)
		}
	}
}

func validateListParams(req *http.Request) (
	limit int, offset int, names []string, err error) {
	// TODO can merge with filters from books?
	// jsonschema validator https://github.com/xeipuuv/gojsonschema
	params := req.URL.Query()

	names, prs := params["name"]
	if !prs {
		names = nil
	}

	var e error

	limitStr, prs := params["limit"]
	if !prs {
		// Default value
		limitStr = []string{"100"} // TODO move to env
	}
	if len(limitStr) > 1 {
		err = fmt.Errorf("%s\r\n%s", err,
			"Only one parameter is valid to 'limit'")
	} else {
		limit, e = strconv.Atoi(limitStr[0])
		if e != nil {
			err = fmt.Errorf("%s\r\n%s", err, e.Error())
		} else if limit < 0 {
			err = fmt.Errorf("%s\r\n%s", err,
				"'limit' parameter must be positive")
			limit = *new(int)
		}
	}

	offsetStr, prs := params["offset"]
	if !prs {
		// Default value
		offsetStr = []string{"0"} // TODO env
	}
	if len(offsetStr) > 1 {
		err = fmt.Errorf("%s\r\n%s", err,
			"Only one parameter is valid to 'offset'")
	} else {
		offset, e = strconv.Atoi(offsetStr[0])
		if e != nil {
			err = fmt.Errorf("%s\r\n%s", err, e.Error())
		} else if offset < 0 {
			err = fmt.Errorf("%s\r\n%s", err,
				"'offset' parameter must be positive")
			offset = *new(int)
		}
	}

	return
}

func getAuthorByID(a *business.Author) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"getAuthor\" route")

		params := mux.Vars(req)
		id, _ := strconv.Atoi(params["id"])

		author, err := a.GetByID(id)

		switch {
		case err != nil:
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")
			// TODO log err
			json.NewEncoder(w).Encode(response{"Fail connection on DB."})
		case author == nil:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(author)
		}
	}
}
