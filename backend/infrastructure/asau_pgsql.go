package infrastructure

import (
	"backend/domain/repository"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DatabaseRepositoryPGSQL struct {
	Database repository.DatabaseInterface
	apiUser  repository.UserRepositoryInterface
}

func NewDatabaseRepositoryPGSQL(
	db repository.DatabaseInterface,
) repository.DatabaseRepositoryPGSQLInterface {
	return &DatabaseRepositoryPGSQL{
		Database: db,
		apiUser:  NewUserRepository(),
	}
}

func (p *DatabaseRepositoryPGSQL) Connect() error {
	log.Printf("DatabaseRepositoryPGSQL open connection")
	return p.Database.Connect()
}

func (p *DatabaseRepositoryPGSQL) Begin() (*sql.Tx, error) {
	log.Printf("DatabaseRepositoryPGSQL begin transaction")
	return p.Database.Begin()
}

func (p *DatabaseRepositoryPGSQL) Close() {
	log.Printf("DatabaseRepositoryPGSQL close connection")
	p.Database.Close()
}

func (p *DatabaseRepositoryPGSQL) Commit(tx *sql.Tx) error {
	log.Printf("DatabaseRepositoryPGSQL commit transaction")
	return p.Database.Commit(tx)
}

func (p *DatabaseRepositoryPGSQL) Rollback(tx *sql.Tx) error {
	log.Printf("DatabaseRepositoryPGSQL rollback transaction")
	return p.Database.Rollback(tx)
}

func (p *DatabaseRepositoryPGSQL) ApiUser() repository.UserRepositoryInterface {
	return p.apiUser
}
