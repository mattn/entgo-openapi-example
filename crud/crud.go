package crud

import (
	"log"

	"github.com/mattn/entgo-openapi-example/ent"
	_ "github.com/mattn/go-sqlite3"
)

var defaultClient *ent.Client

func NewClient() *ent.Client {
	client, err := ent.Open("sqlite3", "file:entry.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	return client
}

func DefaultClient() *ent.Client {
	if defaultClient == nil {
		defaultClient = NewClient()
	}
	return defaultClient
}
