package model

import "time"

type Area struct {
	IDKhuVuc    string     `json:"id_khu_vuc"`
	TenKhuVuc   *string    `json:"ten_khu_vuc"`
	DiaChi      *string    `json:"dia_chi"`
	Create_at   *time.Time `json:"create_at"`
	Update_at   *time.Time `json:"update_at"`
	Deleted_flg *string    `json:"deleted_flg"`
}
