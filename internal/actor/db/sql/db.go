package db

import (
	"context"
	"database/sql"
)

type DB interface {
	GetName() string
	Connect() DB
	Disconnect()
	GetDB() SqlDB
	Ping() error
}

type SqlDB interface {
	Begin() (SqlTx, error)
	DBTX
}

type SqlTx interface {
	Commit() error
	Rollback() error
	DBTX
}

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func newSqlDB(db *sql.DB) *sqlDB {
	return &sqlDB{
		DB: db,
	}
}

type sqlDB struct {
	*sql.DB
}

func (r *sqlDB) Begin() (SqlTx, error) {
	return r.DB.Begin()
}
