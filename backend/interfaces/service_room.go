package interfaces

import (
	"backend/domain/model"
	"backend/domain/repository"
	"database/sql"
)

type RoomServiceInterface interface {
	GetRooms(tx *sql.Tx) ([]*model.Room, error)
	GetRoom(tx *sql.Tx, id_phong string) (*model.Room, error)
	CreateRoom(tx *sql.Tx, room model.Room) error
	UpdateRoom(tx *sql.Tx, room model.Room, id_phong string) error
	DeleteRoom(tx *sql.Tx, id_phong string) error
	DatabaseRepository() repository.DatabaseRepositoryPGSQLInterface
}
