package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// TODO: flag with .csv file name
	// https://golang.org/doc/effective_go.html#init
	// https://golang.org/pkg/flag/#pkg-examples
	db := make(map[int]string)

	csvFile, err := os.Open("Authors.csv")
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for idx, line := range csvLines {
		db[idx] = line[0]
	}
	delete(db, 0)

	fmt.Println(db)
}
