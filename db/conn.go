package db

import (
	"fmt"
	"os"

	"database/sql"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

const file = "file://sponsor-inspector.db"

type Connection struct {
	db *sql.DB
}

func GetConnection() *Connection {
	db, err := sql.Open("libsql", file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", file, err)
		os.Exit(1)
	}

	return &Connection{db: db}
}
