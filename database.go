package gooseneck

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func init() {
	InitLog()
}

type Database struct {
	handle  *sql.DB
	Options *DatabaseOptions
}

type DatabaseOptions struct {
	DSN               string
	MaxOpenConns      int
	MaxIdleConns      int
	ConnMaxLifetime   time.Duration
	ReconnectInterval time.Duration
}

func NewDatabase(options *DatabaseOptions) *Database {
	if options == nil {
		dsn := os.Getenv(DATABASE_URL)
		if dsn == "" {
			Warn().Str("key", DATABASE_URL).Msg("not set")
		}
		options = &DatabaseOptions{
			DSN:               dsn,
			MaxOpenConns:      5,
			MaxIdleConns:      5,
			ConnMaxLifetime:   time.Minute * 5,
			ReconnectInterval: time.Second * 1,
		}
	}
	if options.DSN == "" {
		Fatal().Msg("dsn empty")
	}
	tool := &Database{Options: options}
	tool.connect()
	return tool
}

func (db *Database) connect() {
	if c, err := sql.Open("postgres", db.Options.DSN); err != nil {
		Warn().Err(err).Msg("database connection failed")
		db.handle = &sql.DB{}
	} else {
		c.SetMaxOpenConns(db.Options.MaxOpenConns)
		c.SetMaxIdleConns(db.Options.MaxIdleConns)
		c.SetConnMaxLifetime(db.Options.ConnMaxLifetime)
		db.handle = c
	}
}

func (db *Database) Connection() *sql.DB {
	err := db.handle.Ping()
	count := 0
	for err != nil {
		Info().Msg("database connection lost")
		db.handle.Close()
		time.Sleep(db.Options.ReconnectInterval)
		db.connect()
		count += 1
		Info().Int("count", count).Msg("database reconnection attempt")
		err = db.handle.Ping()
	}
	return db.handle
}
