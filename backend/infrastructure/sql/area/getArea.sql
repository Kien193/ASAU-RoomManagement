SELECT *
FROM khuvuc
WHERE idkhuvuc = $1
  AND deleted_flg = '0'