SELECT *
FROM khachhang
WHERE idkhachhang = $1
  AND deleted_flg = '0'