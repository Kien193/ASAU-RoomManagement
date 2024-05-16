UPDATE khuvuc
SET tenkhuvuc = $2,
  diachi = $3,
  updated_at = NOW()
WHERE idkhuvuc = $1
  AND deleted_flg = '0';