package database

type Book struct {
	ID              int
	Name            string `json:"name"`
	Edition         int    `json:"edition"`
	PublicationYear int    `json:"publicationYear"`
	Authors         []int  `json:"authors"`
}
