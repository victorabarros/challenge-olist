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
	conn, err := sqlx.ConnectContext(ctx, driver, dsn)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &Database{Connection: conn}, nil
}

// StatusCheck returns nil if it can successfully talk to the database.
func (db *Database) StatusCheck() error {
	const q = `SELECT true;`

	if err := db.Connection.Select(&[]bool{}, q); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
