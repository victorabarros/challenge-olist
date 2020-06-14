package database

import (
	"encoding/csv"
	"os"
)

type author struct {
	Name string `json:"name"`
}

// Authors is a abstraction t database Authors
type Authors map[int]author

// LoadCsv to authors
func (a *Authors) LoadCsv(name string) {
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
		(*a)[idx] = author{Name: line[0]}
	}
	delete(*a, 0) // First row is header
}

// SubSection return a sub section
func (a *Authors) SubSection(offset, limit int) Authors { // TODO improve name
	resp := make(Authors, limit)
	for idx := 0; idx < limit; idx++ {
		val, ok := (*a)[idx+offset+1] // `+1` because the firt author id is 1
		if !ok {
			break
		}
		resp[idx+offset+1] = val
	}
	return resp
}

// GetByName return simgle Author
func (a *Authors) GetByName(name string) (Authors, bool) {
	for idx, author := range *a {
		// TODO find a way to return match similars, not only equal
		if author.Name == name {
			return Authors{idx: author}, true
		}
	}

	return nil, false
}

// GetByID return simgle Author
func (a *Authors) GetByID(id int) (Authors, bool) {
	author, prs := (*a)[id]
	if !prs {
		return nil, false
	}

	return Authors{id: author}, true
}

// // NewAuthorsDb return new Authors{}
// func NewAuthorsDb() Authors {
// 	return Authors{}
// }
