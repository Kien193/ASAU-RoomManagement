package service

import (
	"backend/domain/model"
	"backend/domain/repository"
	"backend/interfaces"
	"database/sql"
)

type RoomService struct {
	DatabaseRepositoryPGSQL repository.DatabaseRepositoryPGSQLInterface
}

func NewRoomService(
	DatabaseRepositoryPGSQL repository.DatabaseRepositoryPGSQLInterface,
) interfaces.RoomServiceInterface {
	return &RoomService{
		DatabaseRepositoryPGSQL: DatabaseRepositoryPGSQL,
	}
}

func (service *RoomService) GetRooms(tx *sql.Tx) ([]*model.Room, error) {
	rooms, err := service.DatabaseRepositoryPGSQL.ApiRoom().GetRooms(tx)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (service *RoomService) GetRoom(tx *sql.Tx, id_phong string) (*model.Room, error) {
	room, err := service.DatabaseRepositoryPGSQL.ApiRoom().GetRoom(tx, id_phong)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (service *RoomService) CreateRoom(tx *sql.Tx, room model.Room) error {
	err := service.DatabaseRepositoryPGSQL.ApiRoom().CreateRoom(tx, room)
	if err != nil {
		return err
	}
	return nil
}

func (service *RoomService) UpdateRoom(tx *sql.Tx, room model.Room, id_phong string) error {
	err := service.DatabaseRepositoryPGSQL.ApiRoom().UpdateRoom(tx, room, id_phong)
	if err != nil {
		return err
	}
	return nil
}

func (service *RoomService) DeleteRoom(tx *sql.Tx, id_phong string) error {
	err := service.DatabaseRepositoryPGSQL.ApiRoom().DeleteRoom(tx, id_phong)
	if err != nil {
		return err
	}
	return nil
}

func (service *RoomService) DatabaseRepository() repository.DatabaseRepositoryPGSQLInterface {
	return service.DatabaseRepositoryPGSQL
}
