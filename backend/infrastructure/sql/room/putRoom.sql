UPDATE phong
SET idkhuvuc = $2,
  tenphong = $3,
  sltoida = $4,
  slhientai = $5,
  trangthai = $6,
  trangthaithue = $7,
  giatien = $8,
  updated_at = NOW()
WHERE idphong = $1
  AND deleted_flg = '0';