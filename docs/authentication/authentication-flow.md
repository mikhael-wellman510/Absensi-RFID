docs auth:

Login super admin :
"email": "superadmin@attendance.com",
"password" : "SuperSecret123!"

kenapa dia pakai user session , karena ia menjadi sumber kebenaran
(source of truth) ,
karena supaya bisa tau apakah user sudah logout atau belum , token sudah d cabut,dll

Alur nya adalah :
Login -> access token & refresh token
gunakan acess token untuk hit api
jika access token exp nya berakhir -> maka hit api Refresh token
setelah hit api refresh token , gunakan refresh token yg baru-
supaya menghasilkan access token yg baru

Login -> Access token
Refresh token -> new Access Token & new refresh token