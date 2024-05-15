package service

import (
	"backend/domain/model"
	"backend/domain/repository"
	"backend/interfaces"
	"database/sql"
)

type AreaService struct {
	DatabaseRepositoryPGSQL repository.DatabaseRepositoryPGSQLInterface
}

func NewAreaService(
	DatabaseRepositoryPGSQL repository.DatabaseRepositoryPGSQLInterface,
) interfaces.AreaServiceInterface {
	return &AreaService{
		DatabaseRepositoryPGSQL: DatabaseRepositoryPGSQL,
	}
}

func (service *AreaService) GetAreas(tx *sql.Tx) ([]*model.Area, error) {
	areas, err := service.DatabaseRepositoryPGSQL.ApiArea().GetAreas(tx)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

func (service *AreaService) GetArea(tx *sql.Tx, id_khu_vuc string) (*model.Area, error) {
	area, err := service.DatabaseRepositoryPGSQL.ApiArea().GetArea(tx, id_khu_vuc)
	if err != nil {
		return nil, err
	}
	return area, nil
}

func (service *AreaService) CreateArea(tx *sql.Tx, area model.Area) error {
	err := service.DatabaseRepositoryPGSQL.ApiArea().CreateArea(tx, area)
	if err != nil {
		return err
	}
	return nil
}

func (service *AreaService) UpdateArea(tx *sql.Tx, area model.Area, id_khu_vuc string) error {
	err := service.DatabaseRepositoryPGSQL.ApiArea().UpdateArea(tx, area, id_khu_vuc)
	if err != nil {
		return err
	}
	return nil
}

func (service *AreaService) DeleteArea(tx *sql.Tx, id_khu_vuc string) error {
	err := service.DatabaseRepositoryPGSQL.ApiArea().DeleteArea(tx, id_khu_vuc)
	if err != nil {
		return err
	}
	return nil
}

func (service *AreaService) DatabaseRepository() repository.DatabaseRepositoryPGSQLInterface {
	return service.DatabaseRepositoryPGSQL
}
