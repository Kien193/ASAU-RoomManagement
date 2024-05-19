package model

import "time"

type Room struct {
	IDPhong       string     `json:"id_phong"`
	IDKhuVuc      string     `json:"id_khu_vuc"`
	TenPhong      string     `json:"ten_phong"`
	SLToiDa       int        `json:"sl_toi_da"`
	SLHienTai     int        `json:"sl_hien_tai"`
	TrangThai     int        `json:"trang_thai"`
	TrangThaiThue int        `json:"trang_thai_thue"`
	GiaTien       int        `json:"gia_tien"`
	Create_at     *time.Time `json:"create_at"`
	Update_at     *time.Time `json:"update_at"`
	Deleted_flg   *string    `json:"deleted_flg"`
}
