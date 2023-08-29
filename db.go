package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	DATABASE_URL = "DATABASE_URL"
)

type DatabaseTool struct {
	connection *sql.DB
	dsn        string
}

func NewDatabaseTool() *DatabaseTool {
	dsn := os.Getenv(DATABASE_URL)
	if dsn == "" {
		log.Fatalf("Fatal: environment key %s not set", DATABASE_URL)

	}
	tool := DatabaseTool{dsn: dsn}
	if conn, err := tool.Connect(); err != nil {
		log.Fatalf("Fatal: no db connection.  Error: %s", err)
	} else {
		tool.connection = conn
	}
	return &tool
}

func (db *DatabaseTool) Connect() (*sql.DB, error) {
	if db.connection == nil {
		conn, err := sql.Open("postgres", db.dsn)
		if err != nil {
			return nil, err
		}
		db.connection = conn
		log.Println("Database connected")
	}
	return db.connection, nil
}
