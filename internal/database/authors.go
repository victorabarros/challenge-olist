package database

import (
	"encoding/csv"
	"os"
)

type author struct {
	Name string `json:"name"`
}

type Authors map[int]author

func (a Authors) LoadCsv(name string) {
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
	delete(a, 0) // First row is header
}

func (a Authors) SubSection(offset, limit int) Authors { // TODO improve name
	resp := make(Authors, limit)
	for idx := 0; idx < limit; idx++ {
		val, ok := a[idx+offset+1] // `+1` because the firt author id is 1
		if !ok {
			break
		}
		resp[idx+offset+1] = val
	}
	return resp
}

// // NewAuthorsDb return new Authors{}
// func NewAuthorsDb() Authors {
// 	return Authors{}
// }
