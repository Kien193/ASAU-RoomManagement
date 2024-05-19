INSERT INTO phong (
    idphong,
    idkhuvuc,
    tenphong,
    sltoida,
    slhientai,
    trangthai,
    trangthaithue,
    giatien,
    created_at
  )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    now()
  )