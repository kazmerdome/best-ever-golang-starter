package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type postgresDB struct {
	name          string
	uri           string
	isSslDisabled bool
	db            *sql.DB
	logger        zerolog.Logger
}

func NewPostgresDB(name, uri string, isSslDisabled bool) *postgresDB {
	return &postgresDB{
		name:          name,
		uri:           uri,
		isSslDisabled: isSslDisabled,
		logger: log.
			With().
			Str("actor", "db/postgres").
			Logger(),
	}
}

func (r *postgresDB) GetName() string {
	return r.name
}

func (r *postgresDB) Connect() DB {
	r.logger.
		Info().
		Str("dbName", r.name).
		Str("status", "connecting...").
		Send()

	if r.name == "" {
		r.logger.
			Fatal().
			Str("status", "missing parameters").
			Str("reason", "name is required").
			Send()
	}
	if r.uri == "" {
		r.logger.
			Fatal().
			Str("dbName", r.name).
			Str("status", "missing parameters").
			Str("reason", "uri is required").
			Send()
	}

	connectionUri := fmt.Sprintf("%s/%s", r.uri, r.name)
	if r.isSslDisabled {
		connectionUri = fmt.Sprintf("%s?sslmode=disable", connectionUri)
	}

	db, err := sql.Open("postgres", connectionUri)
	if err != nil {
		r.logger.
			Fatal().
			Err(err).
			Str("status", "connection failed!").
			Send()
	}

	if err = db.Ping(); err != nil {
		r.logger.
			Fatal().
			Err(err).
			Str("status", "failed to ping connection!").
			Send()
	}
	r.db = db
	r.logger.
		Info().
		Str("dbName", r.name).
		Str("status", "connected successfully.").
		Send()
	return r
}

func (r *postgresDB) Disconnect() {
	r.logger.
		Info().
		Str("dbName", r.name).
		Str("status", "disconnecting...").
		Send()

	err := r.db.Close()
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("dbName", r.name).
			Str("status", "failed to disconnect gracefully").
			Send()
	}

	r.logger.
		Info().
		Str("dbName", r.name).
		Str("status", "disconnected successfully").
		Send()
}

func (r *postgresDB) GetDB() SqlDB {
	return newSqlDB(r.db)
}

func (r *postgresDB) Ping() error {
	return r.db.Ping()
}
