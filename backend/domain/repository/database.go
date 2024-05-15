package repository

import "database/sql"

type DatabaseInterface interface {
	Connect() error
	Close()
	Begin() (*sql.Tx, error)
	Rollback(tx *sql.Tx) error
	Commit(tx *sql.Tx) error
}
