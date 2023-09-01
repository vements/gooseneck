package gooseneck

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func init() {
	InitLog()
}

type Database struct {
	Connection *sql.DB
	DSN        string
}

type DatabaseOptions struct {
	DSN          string
	MaxOpenConns int32
}

func NewDatabase(options *DatabaseOptions) *Database {
	if options == nil {
		dsn := os.Getenv(DATABASE_URL)
		if dsn == "" {
			Warn().Str("key", DATABASE_URL).Msg("not set")
		}
		options = &DatabaseOptions{
			DSN:          dsn,
			MaxOpenConns: 5,
		}
	}
	if options.DSN == "" {
		Fatal().Msg("dsn empty")
	}
	tool := Database{DSN: options.DSN}
	if conn, err := tool.Connect(); err != nil {
		Fatal().Err(err).Msg("connection failed")
	} else {
		tool.Connection.SetMaxOpenConns(int(options.MaxOpenConns))
		tool.Connection = conn
	}
	return &tool
}

func (db *Database) Connect() (*sql.DB, error) {
	if db.Connection == nil {
		conn, err := sql.Open("postgres", db.DSN)
		if err != nil {
			return nil, err
		}
		db.Connection = conn
		Info().Msg("database connected")
	}
	return db.Connection, db.Connection.Ping()
}
