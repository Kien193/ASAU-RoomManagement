package service

import (
	"backend/domain/model"
	"backend/domain/repository"
	"backend/interfaces"
	"database/sql"
)

type UserService struct {
	DatabaseRepositoryPGSQL repository.DatabaseRepositoryPGSQLInterface
}

func NewUserService(
	DatabaseRepositoryPGSQL repository.DatabaseRepositoryPGSQLInterface,
) interfaces.UserServiceInterface {
	return &UserService{
		DatabaseRepositoryPGSQL: DatabaseRepositoryPGSQL,
	}
}

func (service *UserService) GetUsers(tx *sql.Tx) ([]*model.User, error) {
	users, err := service.DatabaseRepositoryPGSQL.ApiUser().GetUsers(tx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service *UserService) GetUser(tx *sql.Tx, id_khach_hang string) (*model.User, error) {
	user, err := service.DatabaseRepositoryPGSQL.ApiUser().GetUser(tx, id_khach_hang)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) CreateUser(tx *sql.Tx, user model.User) error {
	err := service.DatabaseRepositoryPGSQL.ApiUser().CreateUser(tx, user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateUser(tx *sql.Tx, user model.User, id_khach_hang string) error {
	err := service.DatabaseRepositoryPGSQL.ApiUser().UpdateUser(tx, user, id_khach_hang)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) DeleteUser(tx *sql.Tx, id_khach_hang string) error {
	err := service.DatabaseRepositoryPGSQL.ApiUser().DeleteUser(tx, id_khach_hang)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) DatabaseRepository() repository.DatabaseRepositoryPGSQLInterface {
	return service.DatabaseRepositoryPGSQL
}
