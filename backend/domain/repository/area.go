package repository

import (
	"backend/domain/model"
	"database/sql"
)

type AreaRepositoryInterface interface {
	GetAreas(tx *sql.Tx) ([]*model.Area, error)
	GetArea(tx *sql.Tx, id_khu_vuc string) (*model.Area, error)
	CreateArea(tx *sql.Tx, area model.Area) error
	UpdateArea(tx *sql.Tx, area model.Area, id_khu_vuc string) error
	DeleteArea(tx *sql.Tx, id_khu_vuc string) error
}
