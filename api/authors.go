package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/victorabarros/challenge-olist/business"
)

type response struct {
	Message string `json:"message"`
}

// listAuthors return list
func listAuthors(a *business.Author) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"listAuthors\" route")

		limit, offset, names, err := validateListParams(req)
		if err != nil {
			// TODO log err
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).
				Encode(response{err.Error()}); err != nil {
				fmt.Println(err)
			}
			return
		}

		authors, err := a.List(limit, offset, names)

		switch {
		case err != nil:
			// TODO log err
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).
				Encode(response{"Fail to query no DB."}); err != nil {
				fmt.Println(err)
			}
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
			err = fmt.Errorf("%s\r\n%s", err,
				e)
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
			err = fmt.Errorf("%s\r\n%s", err, e)
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
			// TODO build error response message
		case author == nil:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(author); err != nil {
				fmt.Println(err)
			}
		}
	}
}
