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

func listAuthors(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json") // TODO Isn't work: >>> requests.get("http://localhost:8092/authors").headers
	if err := json.NewEncoder(w).Encode(db); err != nil {
		fmt.Println(err)
	}
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
	w.Header().Set("Content-Type", "application/json") // TODO Isn't work: {'Date': 'Sun, 14 Jun 2020 00:07:29 GMT', 'Content-Length': '924', 'Content-Type': 'text/plain; charset=utf-8'}
	if err := json.NewEncoder(w).Encode(author); err != nil {
		fmt.Println(err)
	}
}
