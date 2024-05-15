package interfaces

import (
	"backend/domain/model"
	"backend/domain/repository"
	"database/sql"
)

type AreaServiceInterface interface {
	GetAreas(tx *sql.Tx) ([]*model.Area, error)
	GetArea(tx *sql.Tx, id_khu_vuc string) (*model.Area, error)
	CreateArea(tx *sql.Tx, area model.Area) error
	UpdateArea(tx *sql.Tx, area model.Area, id_khu_vuc string) error
	DeleteArea(tx *sql.Tx, id_khu_vuc string) error
	DatabaseRepository() repository.DatabaseRepositoryPGSQLInterface
}
