package main

import (
	"flag"
	"net/http"

	"github.com/victorabarros/work-at-olist/api"
	"github.com/victorabarros/work-at-olist/internal/database"

	"github.com/gorilla/mux"
)

var (
	csvName = flag.String("csv", "Authors.csv", "Authors file")
)

// func init() {
// 	// https://golang.org/doc/effective_go.html#init
// 	// https://github.com/GoogleCloudPlatform/microservices-demo/blob/c78fd12a526c8ba889283ffdbbe4e7d011529935/src/productcatalogservice/server.go#L59
// }

func main() {
	flag.Parse() // `go run main.go -h` for help flag

	db := database.Authors{}
	db.LoadCsv(*csvName)

	srv := newServer(db)
	panic(srv.ListenAndServe())
}

func newServer(db database.Authors) *http.Server {
	r := mux.NewRouter()
	api.SetUpRoutes(r, &db)

	srv := &http.Server{
		Addr:    ":8092", // TODO move to env
		Handler: r,
	}

	return srv
}
