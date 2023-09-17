package db

import (
	"fmt"
	"os"

	"database/sql"

	_ "github.com/libsql/libsql-client-go/libsql"
	migrate "github.com/rubenv/sql-migrate"
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
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	n, err := migrate.Exec(c.db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("failed to apply migrations: %s", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
