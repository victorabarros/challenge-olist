package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/victorabarros/work-at-olist/api"
	"github.com/victorabarros/work-at-olist/internal/database"
)

var (
	csvName = flag.String("csv", "Authors.csv", "Authors file")
	port    = flag.String("port", "8092", "Server port") // TODO move to env
)

// func init() {
// 	// https://golang.org/doc/effective_go.html#init
// 	// https://github.com/GoogleCloudPlatform/microservices-demo/blob/c78fd12a526c8ba889283ffdbbe4e7d011529935/src/productcatalogservice/server.go#L59
// }

func main() {
	flag.Parse() // `go run main.go -h` for help flag

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"olist",
		"1234",
		"127.0.0.1:8093",
		"olist")

	if err != nil {
		panic(err)
	}

	db, err := database.NewDatabase(ctx, "mysql", dsn)
	defer db.Connection.Close()

	db.LoadCsv(*csvName)

	srv := newServer(db)
	fmt.Printf("Up apllication at port %s\n" + *port)
	panic(srv.ListenAndServe())
}

func newServer(db *database.Database) *http.Server {
	r := mux.NewRouter()
	api.SetUpRoutes(r, db, myDb)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
		Handler: r,
	}

	return srv
}
