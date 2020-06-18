package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/victorabarros/challenge-olist/internal/database"
)

// listAuthors return with offset (default = 0) and limit (default = 10) query params
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
			resp, err := db.GetByName(name)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				// TODO add message response
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).
				Encode(resp); err != nil {
				fmt.Println(err)
			}
			return
		}

		authors, err := db.SubSection(limit, offset)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			// TODO add message response
			return

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).
			Encode(authors); err != nil {
			fmt.Println(err)
		}
	}
}

func validateListQueryParams(req *http.Request) (offset int, limit int, name string, errors []error) {
	// TODO improve name
	// jsonschema validator https://github.com/xeipuuv/gojsonschema
	params := req.URL.Query()

	limitStr, prs := params["limit"]
	if !prs {
		limitStr = []string{"100"} // TODO move to env
	}
	// TODO if len > 1 append err
	limit, err := strconv.Atoi(limitStr[0])
	if err != nil {
		errors = append(errors, err)
	}

	offsetStr, prs := params["offset"]
	if !prs {
		offsetStr = []string{"0"}
	}
	// TODO if len > 1 append err
	offset, err = strconv.Atoi(offsetStr[0])
	if err != nil {
		errors = append(errors, err)
	}

	nameQueue, prs := params["name"]
	if !prs {
		return
	}
	// TODO if len > 1 append err? or return more then one?
	name = nameQueue[0]
	return
}

func getAuthor(db *database.Authors) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Starting \"getAuthor\" route")
		params := mux.Vars(req)
		id, _ := strconv.Atoi(params["id"])

		resp, prs := db.GetByID(id)
		if !prs {
			w.WriteHeader(http.StatusNotFound)
			// TODO add message response
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		// TODO Isn't work:
		// >>> requests.get("http://localhost:8092/authors").headers
		// {'Date': 'Sun, 14 Jun 2020 00:07:29 GMT', 'Content-Length': '924', 'Content-Type': 'text/plain; charset=utf-8'}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Println(err)
		}
	}
}
