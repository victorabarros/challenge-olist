package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type author struct {
	Name string `json:"name"`
}

type authors map[int]author

var (
	db      = authors{}
	csvName = flag.String("csv", "Authors.csv", "Authors file")
)

// func init() {
// 	// https://golang.org/doc/effective_go.html#init
// 	// https://github.com/GoogleCloudPlatform/microservices-demo/blob/c78fd12a526c8ba889283ffdbbe4e7d011529935/src/productcatalogservice/server.go#L59
// }

func main() {
	flag.Parse() // `go run main.go -h` for help flag

	db.loadCsv(*csvName)

	srv := newServer()
	panic(srv.ListenAndServe())
}

func (a authors) loadCsv(name string) {
	csvFile, err := os.Open(name)
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic(err)
	}

	for idx, line := range csvLines {
		a[idx] = author{Name: line[0]}
	}
	delete(db, 0) // First row is header
}

func (a authors) subSection(offset, limit int) authors { // TODO improve name
	resp := make(authors, limit)
	for idx := 0; idx < limit; idx++ {
		val, ok := a[idx+offset+1] // `+1` because the firt author id is 1
		if !ok {
			break
		}
		resp[idx+offset+1] = val
	}
	return resp
}

func newServer() *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/",
		func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("work-at-olist\n"))
		}).Methods("GET")

	r.HandleFunc("/authors", listAuthors).Methods("GET")
	r.HandleFunc("/authors/{id:[0-9]+}", getAuthor).Methods("GET")

	srv := &http.Server{
		Addr:    ":8092",
		Handler: r,
	}
	return srv
}

func listAuthors(w http.ResponseWriter, req *http.Request) {
	offset, limit, err := validateListQueryParams(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct{ Message error }{err})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(db.subSection(offset, limit)); err != nil {
		fmt.Println(err)
	}
}

func validateListQueryParams(req *http.Request) (int, int, error) {
	// jsonschema validator https://github.com/xeipuuv/gojsonschema
	params := req.URL.Query()
	limit, ok := params["limit"]
	if !ok {
		limit = []string{"10"} // move to env
	}

	offset, ok := params["offset"]
	if !ok {
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

func getAuthor(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	author, ok := db[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// TODO Isn't work:
	// >>> requests.get("http://localhost:8092/authors").headers
	// {'Date': 'Sun, 14 Jun 2020 00:07:29 GMT', 'Content-Length': '924', 'Content-Type': 'text/plain; charset=utf-8'}
	if err := json.NewEncoder(w).Encode(author); err != nil {
		fmt.Println(err)
	}
}
