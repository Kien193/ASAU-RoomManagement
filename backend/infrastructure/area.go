package infrastructure

import (
	"backend/domain/model"
	"backend/domain/repository"
	"database/sql"
)

type AreaRepository struct{}

func NewAreaRepository() repository.AreaRepositoryInterface {
	return &AreaRepository{}
}

func (r *AreaRepository) GetAreas(tx *sql.Tx) ([]*model.Area, error) {
	return nil, nil
}

func (r *AreaRepository) GetArea(tx *sql.Tx, id_khu_vuc string) (*model.Area, error) {
	return nil, nil
}

func (r *AreaRepository) CreateArea(tx *sql.Tx, area model.Area) error {
	return nil
}

func (r *AreaRepository) UpdateArea(tx *sql.Tx, area model.Area, id_khu_vuc string) error {
	return nil
}

func (r *AreaRepository) DeleteArea(tx *sql.Tx, id_khu_vuc string) error {
	return nil
}
