package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/victorabarros/work-at-olist/internal/database"
)

// listAuthors return with offset (default = 0) and limit (default = 10) query params
func listAuthors(db *database.Authors) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		offset, limit, name, errors := validateListQueryParams(req)

		if len(errors) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(struct{ messages []error }{errors})
			return
		}

		if name != "" {
			resp, prs := db.GetByName(name)
			if !prs {
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

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).
			Encode(db.SubSection(offset, limit)); err != nil {
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
	limit, err := strconv.Atoi(limitStr[0])
	if err != nil {
		errors = append(errors, err)
	}

	offsetStr, prs := params["offset"]
	if !prs {
		offsetStr = []string{"0"}
	}
	offset, err = strconv.Atoi(offsetStr[0])
	if err != nil {
		errors = append(errors, err)
	}

	nameQueue, prs := params["name"]
	if !prs {
		return
	}
	name = nameQueue[0]
	return
}

func getAuthor(db *database.Authors) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
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
