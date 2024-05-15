package infrastructure

import (
	"backend/domain/model"
	"backend/domain/repository"
	"database/sql"
	_ "embed"
	"log"
)

var (
	//go:embed sql/getAllUsers.sql
	getAllUsers string
	//go:embed sql/getUser.sql
	getUser string
	//go:embed sql/postUser.sql
	postUser string
	//go:embed sql/putUser.sql
	putUser string
	//go:embed sql/deleteUser.sql
	deleteUser string
)

type UserRepository struct{}

func NewUserRepository() repository.UserRepositoryInterface {
	return &UserRepository{}
}

func (u *UserRepository) GetUsers(tx *sql.Tx) ([]*model.User, error) {
	var users []*model.User
	rows, err := tx.Query(getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.IDKhachHang,
			&user.TenKH,
			&user.SoDienThoai,
			&user.NgaySinh,
			&user.DiaChi,
			&user.GioiTinh,
			&user.Create_at,
			&user.Update_at,
			&user.Deleted_flg,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *UserRepository) GetUser(tx *sql.Tx, id_khach_hang string) (*model.User, error) {
	var user model.User
	err := tx.QueryRow(getUser, id_khach_hang).Scan(
		&user.IDKhachHang,
		&user.TenKH,
		&user.SoDienThoai,
		&user.NgaySinh,
		&user.DiaChi,
		&user.GioiTinh,
		&user.Create_at,
		&user.Update_at,
		&user.Deleted_flg,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) CreateUser(tx *sql.Tx, user model.User) error {
	stmt, err := tx.Prepare(postUser)
	if err != nil {
		log.Println("Post user, prepare error: ", err)
		return err
	}
	_, err = stmt.Exec(
		&user.IDKhachHang,
		&user.TenKH,
		&user.SoDienThoai,
		&user.NgaySinh,
		&user.DiaChi,
		&user.GioiTinh,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) UpdateUser(tx *sql.Tx, user model.User, id_khach_hang string) error {
	stmt, err := tx.Prepare(putUser)
	if err != nil {
		log.Println("Put user, prepare error: ", err)
		return err
	}
	_, err = stmt.Exec(
		&id_khach_hang,
		&user.TenKH,
		&user.SoDienThoai,
		&user.NgaySinh,
		&user.DiaChi,
		&user.GioiTinh,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) DeleteUser(tx *sql.Tx, id_khach_hang string) error {
	stmt, err := tx.Prepare(deleteUser)
	if err != nil {
		log.Println("Delete user, prepare error: ", err)
		return err
	}
	_, err = stmt.Exec(id_khach_hang)
	if err != nil {
		return err
	}
	return nil
}
