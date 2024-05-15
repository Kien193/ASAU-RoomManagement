INSERT INTO khachhang (
    idkhachhang,
    tenkh,
    sodienthoai,
    ngaysinh,
    diachi,
    gioitinh,
    created_at
  )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    now()
  )