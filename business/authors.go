package business

import (
	"encoding/csv"
	"os"
	"sync"

	"github.com/victorabarros/challenge-olist/internal/database"
)

// Author business layer
type Author struct {
	DB database.Database
	wg sync.WaitGroup
}

// LoadCsv load csv authors to database
func (a *Author) LoadCsv(csv string) {
	names := readNames(csv)
	namesDeduplicates := deduplicate(names)

	a.wg.Add(len(namesDeduplicates))
	for _, name := range namesDeduplicates {
		go func(name string) {
			a.DB.CreateAuthor(name)
			a.wg.Done()
		}(name)
	}
	a.wg.Wait()
}

func readNames(csvName string) []string {
	csvFile, err := os.Open(csvName)
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic(err)
	}

	// Only first column import to us
	names := make([]string, len(csvLines)-1)
	for idx, line := range csvLines[1:] {
		names[idx] = line[0]
	}

	return names
}

// TODO works with interface{}?
func deduplicate(arr []string) []string {
	keys := make(map[string]bool)
	ans := []string{}

	for _, val := range arr {
		_, prs := keys[val]
		if !prs {
			ans = append(ans, val)
		}
	}

	return ans
}

// List return authors using params inputed // TODO improve message
func (a *Author) List(offset, limit int, name string) (*[]database.Author, error) {
	if name != "" {
		author, err := a.DB.GetAuthorByName(name)
		if err != nil {
			return nil, err
		}
		return &author, nil
	}

	authors, err := a.DB.ListAuthors(limit, offset)
	if err != nil {
		return nil, err
	}

	return &authors, nil
}
