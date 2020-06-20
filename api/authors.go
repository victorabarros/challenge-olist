package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/victorabarros/challenge-olist/internal/database"
)

// listAuthors return list
func listAuthors(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"listAuthors\" route")
		offset, limit, name, errors := validateListQueryParams(req)

		if len(errors) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(struct{ messages []error }{errors})
			return
		}

		if name != "" {
			author, err := db.GetAuthorByName(name)
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			// TODO se len(author) == 0 http.StatusNotFound

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).
				Encode(author); err != nil {
				fmt.Println(err)
			}
			return
		}

		authors, err := db.ListAuthors(limit, offset)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			// TODO add message response
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).
			Encode(authors); err != nil {
			fmt.Println(err)
		}
	}
}

// TODO improve name
func validateListQueryParams(req *http.Request) (
	offset, limit int, name string, errors []error) {
	// jsonschema validator https://github.com/xeipuuv/gojsonschema
	params := req.URL.Query()

	limitStr, prs := params["limit"]
	if !prs {
		// Default value
		limitStr = []string{"100"} // TODO move to env
	}
	// TODO if len > 1 append err
	limit, err := strconv.Atoi(limitStr[0])
	if err != nil {
		errors = append(errors, err)
	}

	offsetStr, prs := params["offset"]
	if !prs {
		// Default value
		offsetStr = []string{"0"}
	}
	// TODO if len > 1 append err
	offset, err = strconv.Atoi(offsetStr[0])
	if err != nil {
		errors = append(errors, err)
	}

	nameQueue, prs := params["name"]
	if prs {
		// TODO if len > 1 append err? or return more then one?
		name = nameQueue[0]
	}
	return
	// TODO: validar se são offset e limit são inteiros positivos
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
