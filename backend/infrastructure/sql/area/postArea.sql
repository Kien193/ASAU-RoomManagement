INSERT INTO khachhang (
    idkhuvuc,
    tenkhuvuc,
    diachi,
    created_at
  )
VALUES ($1, $2, $3, NOW())