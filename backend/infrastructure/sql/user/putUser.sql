UPDATE khachhang
SET tenkh = $2,
  sodienthoai = $3,
  ngaysinh = $4,
  diachi = $5,
  gioitinh = $6,
  updated_at = NOW()
WHERE idkhachhang = $1
  AND deleted_flg = '0';