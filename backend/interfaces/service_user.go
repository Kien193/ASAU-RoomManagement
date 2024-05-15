package interfaces

import (
	"backend/domain/model"
	"backend/domain/repository"
	"database/sql"
)

type UserServiceInterface interface {
	GetUsers(tx *sql.Tx) ([]*model.User, error)
	GetUser(tx *sql.Tx, id_khach_hang string) (*model.User, error)
	CreateUser(tx *sql.Tx, user model.User) error
	UpdateUser(tx *sql.Tx, user model.User, id_khach_hang string) error
	DeleteUser(tx *sql.Tx, id_khach_hang string) error
	DatabaseRepository() repository.DatabaseRepositoryPGSQLInterface
}
