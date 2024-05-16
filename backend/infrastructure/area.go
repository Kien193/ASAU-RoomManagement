package infrastructure

import (
	"backend/domain/model"
	"backend/domain/repository"
	"database/sql"
	_ "embed"
	"log"
)

var (
	//go:embed sql/area/getAllAreas.sql
	getAllAreas string
	//go:embed sql/area/getArea.sql
	getArea string
	//go:embed sql/area/postArea.sql
	postArea string
	//go:embed sql/area/putArea.sql
	putArea string
	//go:embed sql/area/deleteArea.sql
	deleteArea string
)

type AreaRepository struct{}

func NewAreaRepository() repository.AreaRepositoryInterface {
	return &AreaRepository{}
}

func (r *AreaRepository) GetAreas(tx *sql.Tx) ([]*model.Area, error) {
	var areas []*model.Area
	rows, err := tx.Query(getAllAreas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var area model.Area
		err := rows.Scan(
			&area.IDKhuVuc,
			&area.TenKhuVuc,
			&area.DiaChi,
			&area.Create_at,
			&area.Update_at,
			&area.Deleted_flg,
		)
		if err != nil {
			return nil, err
		}
		areas = append(areas, &area)
	}
	return areas, nil
}

func (r *AreaRepository) GetArea(tx *sql.Tx, id_khu_vuc string) (*model.Area, error) {
	var area model.Area
	err := tx.QueryRow(getArea, id_khu_vuc).Scan(
		&area.IDKhuVuc,
		&area.TenKhuVuc,
		&area.DiaChi,
		&area.Create_at,
		&area.Update_at,
		&area.Deleted_flg,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &area, nil
}

func (r *AreaRepository) CreateArea(tx *sql.Tx, area model.Area) error {
	stmt, err := tx.Prepare(postArea)
	if err != nil {
		log.Println("Post area, prepare error: ", err)
		return err
	}
	_, err = stmt.Exec(
		&area.IDKhuVuc,
		&area.TenKhuVuc,
		&area.DiaChi,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *AreaRepository) UpdateArea(tx *sql.Tx, area model.Area, id_khu_vuc string) error {
	stmt, err := tx.Prepare(putArea)
	if err != nil {
		log.Println("Put area, prepare error: ", err)
		return err
	}
	_, err = stmt.Exec(
		&id_khu_vuc,
		&area.TenKhuVuc,
		&area.DiaChi,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *AreaRepository) DeleteArea(tx *sql.Tx, id_khu_vuc string) error {
	stmt, err := tx.Prepare(deleteArea)
	if err != nil {
		log.Println("Delete area, prepare error: ", err)
		return err
	}
	_, err = stmt.Exec(id_khu_vuc)
	if err != nil {
		return err
	}
	return nil
}
