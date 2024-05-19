SELECT *
FROM phong
WHERE idphong = $1
  AND deleted_flg = '0'