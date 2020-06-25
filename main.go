package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/victorabarros/challenge-olist/api"
	"github.com/victorabarros/challenge-olist/business"
	"github.com/victorabarros/challenge-olist/internal/database"
)

var (
	csvName = flag.String("csv", "Authors.csv", "Authors file")
)

func main() {
	flag.Parse() // `go run main.go -h` for help flag

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// TODO use internal/configuration
	// TODO set loglevel
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"olist",
		"1234",
		"127.0.0.1:8093",
		"olist")

	db, err := database.NewDatabase(ctx, "mysql", dsn)
	defer db.Connection.Close()
	if err != nil {
		panic(err)
	}

	author := business.Author{
		DB: db,
	}

	book := business.Book{
		DB: db,
	}

	author.LoadCsv(*csvName)

	srv := newServer(&author, book)
	fmt.Printf("Up apllication at port %s\n", "8092")
	panic(srv.ListenAndServe())
}

func newServer(a *business.Author, b business.Book) *http.Server {
	r := mux.NewRouter()
	api.SetUpRoutes(r, a, b)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", "8092"),
		Handler: r,
	}

	return &srv
}
