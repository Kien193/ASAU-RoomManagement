package infrastructure

import (
	"backend/domain/model"
	"backend/domain/repository"
	"database/sql"
	_ "embed"
	"log"
)

var (
	//go:embed sql/room/getAllRooms.sql
	getAllRooms string
	//go:embed sql/room/getRoom.sql
	getRoom string
	//go:embed sql/room/postRoom.sql
	postRoom string
	//go:embed sql/room/putRoom.sql
	putRoom string
	//go:embed sql/room/deleteRoom.sql
	deleteRoom string
)

type RoomRepository struct{}

func NewRoomRepository() repository.RoomRepositoryInterface {
	return &RoomRepository{}
}

func (r *RoomRepository) GetRooms(tx *sql.Tx) ([]*model.Room, error) {
	var rooms []*model.Room
	rows, err := tx.Query(getAllRooms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var room model.Room
		err := rows.Scan(
			&room.IDPhong,
			&room.IDKhuVuc,
			&room.TenPhong,
			&room.SLToiDa,
			&room.SLHienTai,
			&room.TrangThai,
			&room.TrangThaiThue,
			&room.Create_at,
			&room.Update_at,
			&room.Deleted_flg,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

func (r *RoomRepository) GetRoom(tx *sql.Tx, id_phong string) (*model.Room, error) {
	var room model.Room
	err := tx.QueryRow(getRoom, id_phong).Scan(
		&room.IDPhong,
		&room.IDKhuVuc,
		&room.TenPhong,
		&room.SLToiDa,
		&room.SLHienTai,
		&room.TrangThai,
		&room.TrangThaiThue,
		&room.Create_at,
		&room.Update_at,
		&room.Deleted_flg,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) CreateRoom(tx *sql.Tx, room model.Room) error {
	stmt, err := tx.Prepare(postRoom)
	if err != nil {
		log.Println("Post room, prepare error: ", err)
		return err
	}

	_, err = stmt.Exec(
		&room.IDPhong,
		&room.IDKhuVuc,
		&room.TenPhong,
		&room.SLToiDa,
		&room.SLHienTai,
		&room.TrangThai,
		&room.TrangThaiThue,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomRepository) UpdateRoom(tx *sql.Tx, room model.Room, id_phong string) error {
	stmt, err := tx.Prepare(putRoom)
	if err != nil {
		log.Println("Put room, prepare error: ", err)
		return err
	}

	_, err = stmt.Exec(
		&id_phong,
		&room.IDKhuVuc,
		&room.TenPhong,
		&room.SLToiDa,
		&room.SLHienTai,
		&room.TrangThai,
		&room.TrangThaiThue,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomRepository) DeleteRoom(tx *sql.Tx, id_phong string) error {
	stmt, err := tx.Prepare(deleteRoom)
	if err != nil {
		log.Println("Delete room, prepare error: ", err)
		return err
	}

	_, err = stmt.Exec(id_phong)
	if err != nil {
		return err
	}
	return nil
}
