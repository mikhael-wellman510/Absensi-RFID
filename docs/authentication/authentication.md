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

1. buat utils untuk middleware -> untuk error handling nya (DONE)
2. buat utils untuk rate+limiters -> (DONE)
3. buat main nya (DONE)
4. Refactor, repo, dan usecase -> user-session , audit-log (DONE)
5. Testing , login multiple (max 3) ,login 2 akun , lalu logout all (DONE)
6. Ganti nama nama yg di hardcode (DONE)
7. Testing tanpa token / token exp , dia gagal (DONE)
8. testing Role (DONE)
9. buatkan transactional
10. ubah semua data yg butuh users , untuk ambil dari auth
11. nnti tambahkan email verification Ketika register
12. Test forgot password dan reset password


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

4. Dia tidak akan bisa membuat Access token baru ketika sudah di revoked (is_revoked=1)
5. Kenapa refresh token nya di hash , karena jika terjadi kebocoran di DB ,ia tidak menyimpan
   refresh token yg asli , sehingga jika ia ingin generate API refresh token , ia harus menebak
   atau gmna cara nya match antara request dan data refresh Token yg di hash di db.
   Jadi , untuk keamanan jika database di jebol
6. Untuk logout all , nnti core nya ada di middleware , ketika dia cari session ID,
   lalu cek apakah revoke nya true(1) ? jika true , maka dia tidak bisa hit API 