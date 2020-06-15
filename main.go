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
)

// func init() {
// 	// https://golang.org/doc/effective_go.html#init
// 	// https://github.com/GoogleCloudPlatform/microservices-demo/blob/c78fd12a526c8ba889283ffdbbe4e7d011529935/src/productcatalogservice/server.go#L59
// }

func main() {
	flag.Parse() // `go run main.go -h` for help flag

	db := database.Authors{}
	db.LoadCsv(*csvName)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("ctx")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"root",
		"my-secret-pw",
		"127.0.0.1:8093",
		"olist")

	fmt.Println("db1")
	myDB, err := database.NewDatabase(ctx, "mysql", dsn)
	fmt.Println("db2")
	defer myDB.Connection.Close()
	if err != nil {
		panic(err)
	}

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
