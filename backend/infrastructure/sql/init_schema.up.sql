CREATE TABLE IF NOT EXISTS KhachHang(
	IDKhachHang varchar(12) PRIMARY KEY,
	TenKH varchar(50),
	SoDienThoai varchar(20),
	NgaySinh varchar(20),
	DiaChi varchar(100),
	GioiTinh varchar(20),
	created_at timestamp,
	updated_at timestamp,
	deleted_flg varchar(1) DEFAULT '0'
);
CREATE TABLE IF NOT EXISTS Phong(
	IDPhong varchar(14) PRIMARY KEY,
	IDKhuVuc varchar(14) NOT NULL,
	TenPhong varchar(50),
	SLToiDa int,
	SLHienTai int,
	TrangThai int,
	TrangThaiThue int,
	GiaTien int,
	created_at timestamp,
	updated_at timestamp,
	deleted_flg varchar(1) DEFAULT '0'
);
CREATE TABLE IF NOT EXISTS KhuVuc(
	IDKhuVuc varchar(14) PRIMARY KEY,
	TenKhuVuc varchar(50),
	DiaChi varchar(100),
	created_at timestamp,
	updated_at timestamp,
	deleted_flg varchar(1) DEFAULT '0'
);
CREATE TABLE IF NOT EXISTS HoaDon(
	IDHoaDon varchar(14) PRIMARY KEY,
	IDKhachHang varchar(12) NOT NULL,
	TongTien int,
	created_at timestamp,
	updated_at timestamp,
	deleted_flg varchar(1) DEFAULT '0'
);
CREATE TABLE IF NOT EXISTS ChiTietHoaDon(
	IDChiTietHoaDon varchar(14) PRIMARY KEY,
	IDHoaDon varchar(14) NOT NULL,
	IDPhong varchar(14) NOT NULL,
	SoKW int,
	SoNuoc int,
	ChiSoNuocMoi int,
	ChiSoNuocCu int,
	ChiSoDienMoi int,
	ChiSoDienCu int,
	created_at timestamp,
	updated_at timestamp,
	deleted_flg varchar(1) DEFAULT '0'
);
CREATE TABLE IF NOT EXISTS HopDong(
	IDHopDong varchar(14) PRIMARY KEY,
	IDKhachHang varchar(12) NOT NULL,
	TienCoc int,
	created_at timestamp,
	updated_at timestamp,
	deleted_flg varchar(1) DEFAULT '0'
);
CREATE TABLE IF NOT EXISTS ChiTietHopDong(
	IDChiTietHopDong varchar(14) PRIMARY KEY,
	IDHopDong varchar(14) NOT NULL,
	IDPhong varchar(14) NOT NULL,
	ThoiHanThue varchar(20),
	NgayBatDau date,
	NgayKetThuc date,
	created_at timestamp,
	updated_at timestamp,
	deleted_flg varchar(1) DEFAULT '0'
);
CREATE TABLE IF NOT EXISTS Admins(
	username varchar(50) DEFAULT 'admin',
	passwords varchar(50) DEFAULT '123456'
);