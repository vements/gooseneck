package gooseneck

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DatabaseTool struct {
	Connection *sql.DB
	DSN        string
}

func NewDatabaseTool() *DatabaseTool {
	dsn := os.Getenv(DATABASE_URL)
	if dsn == "" {
		log.Fatalf("Fatal: environment key %s not set", DATABASE_URL)

	}
	tool := DatabaseTool{DSN: dsn}
	if conn, err := tool.Connect(); err != nil {
		log.Fatalf("Fatal: no db connection.  Error: %s", err)
	} else {
		tool.Connection = conn
	}
	return &tool
}

func (db *DatabaseTool) Connect() (*sql.DB, error) {
	if db.Connection == nil {
		conn, err := sql.Open("postgres", db.DSN)
		if err != nil {
			return nil, err
		}
		db.Connection = conn
		log.Println("Database connected")
	}
	return db.Connection, nil
}
