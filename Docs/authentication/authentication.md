Refresh Token Rotation & Invalidation (OK)

Device/Session Management (Logout & Logout All Devices)

Account Lockout Protection (Anti-Brute Force)

Audit Logging (Pencatatan aktivitas login & security events)

Role-Based Access Control (RBAC) Middleware

Password Reset & Email Verification Flow

Security Hardening & Rate Limiting

pelajari security!!

===============
Deadline

1. buat utils untuk middleware -> untuk error handling nya
2. buat utils untuk rate+limiters
3. ubah semua data yg butuh users , untuk ambil dari auth
4. nnti tambahkan email verification Ketika register
5. buat main nya (done)
6. Refactor, repo, dan usecase -> user-session , audit-log (DONE)
7. Testing , login multiple ,login 2 akun , lalu logout all
8. Ganti nama nama yg di hardcode
9. testing Role
10. Testing tanpa token / token exp , dia gagal (DONE)
11. buatkan transactional


=====================

1. Fungsi ada access token dan refresh token yaitu supaya access token nya tidak mudah di curi ,
   karena kalua di curi , nnti access token nya bisa di pakai oleh orang lain ,
   nah berarti cara nya adalah dengan memperpendek access token nya menjadi 30 detik ,
   sehingga dengan ada nya refresh token , nnti dia akan membuat access token baru
   dan refresh token harus di hash ke DB , supaya tidak di curi

2. JWT (Json web token)
   di dalam nya ada header , payload , dan token . jadi ada informasi nya

3. Stateless dan statefull
   kalau stateless : tidak menyimpan data (Access Token)
   kalau statefull : menyimpan data (Refresh token ke DB)
4. 