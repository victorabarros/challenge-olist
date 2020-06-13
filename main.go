package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

var (
	authors = make(map[int]string)
	csvName = flag.String("csv", "Authors.csv", "Authors file")
)

// func init() {
// 	// https://golang.org/doc/effective_go.html#init
// 	// https://github.com/GoogleCloudPlatform/microservices-demo/blob/c78fd12a526c8ba889283ffdbbe4e7d011529935/src/productcatalogservice/server.go#L59
// }

func main() {
	flag.Parse() // `go run main.go -h` for help flag
	loadAuthors()
	fmt.Println(authors)
}

func loadAuthors() {
	csvFile, err := os.Open(*csvName)
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for idx, line := range csvLines {
		authors[idx] = line[0]
	}
	delete(authors, 0) // First row is header
}
