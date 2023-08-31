package gooseneck

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	Connection *sql.DB
	DSN        string
}

func NewDB() *DB {
	dsn := os.Getenv(DATABASE_URL)
	if dsn == "" {
		Fatal().Str("key", DATABASE_URL).Msg("not set")
	}
	tool := DB{DSN: dsn}
	if conn, err := tool.Connect(); err != nil {
		Fatal().Err(err).Msg("connection failed")
	} else {
		tool.Connection = conn
	}
	return &tool
}

func (db *DB) Connect() (*sql.DB, error) {
	if db.Connection == nil {
		conn, err := sql.Open("postgres", db.DSN)
		if err != nil {
			return nil, err
		}
		db.Connection = conn
		Info().Msg("database connected")
	}
	return db.Connection, nil
}
