package model

import "time"

type User struct {
	IDKhachHang string     `json:"id_khach_hang"`
	TenKH       *string    `json:"ten_kh"`
	SoDienThoai *string    `json:"so_dien_thoai"`
	NgaySinh    *string    `json:"ngay_sinh"`
	DiaChi      *string    `json:"dia_chi"`
	GioiTinh    *string    `json:"gioi_tinh"`
	Create_at   *time.Time `json:"create_at"`
	Update_at   *time.Time `json:"update_at"`
	Deleted_flg *string    `json:"deleted_flg"`
}
