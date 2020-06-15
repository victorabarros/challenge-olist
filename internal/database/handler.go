package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Database struct with connection
type Database struct {
	Connection *sqlx.DB
}

// NewDatabase return database connection.
func NewDatabase(ctx context.Context, driver string, dsn string) (*Database, error) {
	fmt.Println("db11") // TODO remove
	conn, err := sqlx.ConnectContext(ctx, driver, dsn)
	fmt.Println("db12") // TODO remove
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &Database{Connection: conn}, nil
}
