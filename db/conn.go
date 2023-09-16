package db

import (
	"fmt"
	"os"

	"database/sql"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

type Connection struct {
	db *sql.DB
}

func getPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("failed to get working directory: %s", err))
	}
	return fmt.Sprintf("file://%s/%s", wd, "sponsor-inspector.db")
}

func GetConnection() *Connection {
	file := getPath()
	db, err := sql.Open("libsql", file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", file, err)
		os.Exit(1)
	}

	c := &Connection{db: db}
	c.migrate()

	return c
}

func (c *Connection) migrate() {
	file := getPath()
	u, _ := url.Parse(file)
	db := dbmate.New(u)

	if err := db.CreateAndMigrate(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to migrate db %s: %s", file, err)
		os.Exit(1)
	}
}
