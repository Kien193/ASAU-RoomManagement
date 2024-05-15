package repository

import (
	"backend/domain/model"
	"database/sql"
)

type UserRepositoryInterface interface {
	GetUsers(tx *sql.Tx) ([]*model.User, error)
	GetUser(tx *sql.Tx, id_khach_hang string) (*model.User, error)
	CreateUser(tx *sql.Tx, user model.User) error
	UpdateUser(tx *sql.Tx, user model.User, id_khach_hang string) error
	DeleteUser(tx *sql.Tx, id_khach_hang string) error
}
