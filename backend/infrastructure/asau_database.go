package infrastructure

import (
	"backend/domain/repository"
	"database/sql"
	"fmt"
)

type Database struct {
	Dsn             string
	Username        string
	Password        string
	GeneralDatabase *sql.DB
}

func NewDatabase(
	dsn string,
	username string,
	password string,
) repository.DatabaseInterface {
	return &Database{
		Dsn:      dsn,
		Username: username,
		Password: password,
	}
}

func (db *Database) Close() {
	if db.GeneralDatabase != nil {
		db.GeneralDatabase.Close()
		db.GeneralDatabase = nil
	}
}

func (db *Database) Begin() (*sql.Tx, error) {
	txDb, errBegin := db.GeneralDatabase.Begin()
	if errBegin != nil {
		return nil, errBegin
	}
	return txDb, nil
}

func (db *Database) Rollback(tx *sql.Tx) error {
	errRollback := tx.Rollback()
	if errRollback != nil {
		return errRollback
	}
	return nil
}

func (db *Database) Commit(tx *sql.Tx) error {
	errCommit := tx.Commit()
	if errCommit != nil {
		return errCommit
	}
	return nil
}

func (db *Database) Connect() error {
	errConnect := db.ConnectToPgsql()
	if errConnect != nil {
		return errConnect
	}
	return nil
}

func (db *Database) ConnectToPgsql() error {
	var (
		dsn            string
		username       string
		password       string
		errGeneralOpen error
		errGeneralPing error
	)
	dsn = db.Dsn
	username = db.Username
	password = db.Password
	if db.GeneralDatabase == nil {
		dsn = fmt.Sprintf("%s user=%s password=%s", dsn, username, password)
		// fmt.Println(dsn)
		db.GeneralDatabase, errGeneralOpen = sql.Open("postgres", dsn)
		if errGeneralOpen != nil {
			return errGeneralOpen
		}
		errGeneralPing = db.GeneralDatabase.Ping()
		if errGeneralPing != nil {
			return errGeneralPing
		}
	}
	return nil
}
