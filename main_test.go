package main

import (
	"testing"

	"github.com/victorabarros/challenge-olist/business"
)

func testNewServer(t *testing.T) {
	author := business.Author{}
	book := business.Book{}

	_ = newServer(&author, book)
}
