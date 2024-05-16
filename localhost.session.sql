-- INSERT INTO khachhang (
--     idkhachhang,
--     tenkh,
--     sodienthoai,
--     ngaysinh,
--     diachi,
--     gioitinh,
--     created_at
--   )
-- VALUES (
--     '111111111111',
--     'Tiến Phát',
--     '0911111111',
--     '19-03-2002',
--     '183 Bình Quới',
--     'nam',
--     NOW()
--   );
SELECT *
FROM khachhang
WHERE deleted_flg = '0'