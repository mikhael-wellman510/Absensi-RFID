package usecases

import (
	"attendance-api/config"
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"attendance-api/internal/utils"
	"attendance-api/pkg/constants"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type (
	AuthService interface {
		CreateUser(ctx context.Context, req *entities.CreateUserRequest) (*entities.UserResponse, error)
		Login(ctx context.Context, req *entities.LoginRequest, ip, userAgent string) (*entities.AuthResponse, error)
		RefreshToken(ctx context.Context, req *entities.RefreshTokenRequest, ip, userAgent string) (*entities.AuthResponse, error)
		Logout(ctx context.Context, sessionID string) error
		LogoutAll(ctx context.Context, userID string) error
		GetMe(ctx context.Context, userID string) (*entities.User, error)
		ForgotPassword(ctx context.Context, req *entities.ForgotPasswordRequest, ip, userAgent string) error
		ResetPassword(ctx context.Context, req *entities.ResetPasswordRequest) error
		VerifyEmail(ctx context.Context, token string) error
	}

	authService struct {
		userSessionService UserSessionService
		authRepository     repositories.AuthRepository
		auditLogService    AuditLogService
	}
)

func NewAuthService(authRepository repositories.AuthRepository, auditLogService AuditLogService, userSessionService UserSessionService) AuthService {
	return &authService{
		authRepository:     authRepository,
		auditLogService:    auditLogService,
		userSessionService: userSessionService,
	}
}

func (a *authService) CreateUser(ctx context.Context, req *entities.CreateUserRequest) (*entities.UserResponse, error) {
	// 1. Cek apakah email sudah digunakan
	existingEmailUser, err := a.authRepository.FindByEmail(ctx, req.Email)
	if err == nil && existingEmailUser != nil {
		return nil, errors.New("email sudah terdaftar")
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 2. Cek apakah nomor telepon sudah digunakan
	existingPhoneUser, err := a.authRepository.FindByPhoneNumber(ctx, req.PhoneNumber)
	if err == nil && existingPhoneUser != nil {
		return nil, errors.New("nomor telepon sudah terdaftar")
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 3. Hash Password dengan Bcrypt
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("gagal memproses password")
	}

	// 4. Generate Token Verifikasi Email
	rawVerificationToken, _ := utils.GenerateRandomToken()
	hashedVerificationToken := utils.HashToken(rawVerificationToken)

	// 5. Inisialisasi Entity User Baru
	user := &entities.User{
		FullName:               req.FullName,
		Email:                  req.Email,
		PhoneNumber:            req.PhoneNumber,
		Password:               hashedPassword,
		Role:                   req.Role,
		IsActive:               true,
		IsEmailVerified:        false,
		EmailVerificationToken: hashedVerificationToken,
	}

	if err := a.authRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	// TODO: Panggil EmailService untuk mengirim link verifikasi menggunakan `rawVerificationToken`

	return &entities.UserResponse{
		Id:              user.ID,
		FullName:        user.FullName,
		Email:           user.Email,
		PhoneNumber:     user.PhoneNumber,
		Role:            user.Role,
		IsActive:        user.IsActive,
		IsEmailVerified: user.IsEmailVerified,
		LastLogin:       user.LastLogin,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}, nil
}

func (a *authService) Login(ctx context.Context, req *entities.LoginRequest, ip, userAgent string) (*entities.AuthResponse, error) {
	user, err := a.authRepository.FindByEmail(ctx, req.Email)

	// Cek email ada atau engga
	if err != nil {
		a.auditLogService.logAudit(ctx, nil, constants.LoginFailed, ip, userAgent, fmt.Sprintf("Email tidak ditemukan: %s", req.Email))
		return nil, errors.New("email atau password salah")
	}

	// 1. Cek Account terkunci atau engga
	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		return nil, fmt.Errorf("akun Anda terkunci hingga %s karena terlalu banyak percobaan gagal", user.LockedUntil.Format("15:04:05"))
	}

	// 2. Jika Password salah , masuk kesini
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		user.FailedLoginAttempts++
		maxAttempts, _ := strconv.Atoi(config.Config(constants.MaxLoginAttempts))
		if maxAttempts == 0 {
			maxAttempts = 5
		}

		// Jika percobaan login lebih dari > 5
		if user.FailedLoginAttempts >= maxAttempts {
			user.LockedUntil = new(time.Now().Add(15 * time.Minute))
			a.auditLogService.logAudit(ctx, &user.ID, constants.AccountLocked, ip, userAgent, constants.ToManyRequest)
		}

		// Edit untuk menambahkan failed login attemps
		_ = a.authRepository.UpdateUser(ctx, user)
		a.auditLogService.logAudit(ctx, &user.ID, constants.LoginFailed, ip, userAgent, constants.WrongPassword)
		return nil, errors.New("email atau password salah")
	}

	// 3. Reset percobaan gagal jika sukses
	now := time.Now()
	user.FailedLoginAttempts = 0
	user.LockedUntil = nil
	user.LastLogin = now
	_ = a.authRepository.UpdateUser(ctx, user)

	// 4. Generate Refresh Token & Simpan Session
	rawRefreshToken, _ := utils.GenerateRandomToken()
	log.Println("raw refresh token : ", rawRefreshToken)
	refreshTokenHash := utils.HashToken(rawRefreshToken)
	log.Println("refresh Token Hash : ", refreshTokenHash)

	ttlDays, _ := strconv.Atoi(config.Config(constants.RefreshTokenTtlDays))
	if ttlDays == 0 {
		ttlDays = 7
	}

	// Sesion di simpan , supaya bisa memberikan informasi sesi
	// yg terutama dari refresh token , supaya bisa logout multi device
	session := &entities.UserSession{
		UserID:           user.ID,
		RefreshTokenHash: refreshTokenHash,
		IpAddress:        ip,
		UserAgent:        userAgent,
		ExpiresAt:        time.Now().AddDate(0, 0, ttlDays),
	}

	if err := a.userSessionService.CreateSession(ctx, session); err != nil {
		return nil, err
	}

	// 5. Generate Access Token JWT
	jwtTTL, _ := strconv.Atoi(config.Config(constants.JwtTtlMinutes))
	if jwtTTL == 0 {
		jwtTTL = 15
	}

	accessToken, err := utils.GenerateAccessToken(user.ID, user.Role, session.ID, config.Config(constants.EnvJwtSecret), jwtTTL)
	if err != nil {
		return nil, err
	}

	a.auditLogService.logAudit(ctx, &user.ID, constants.LoginSuccess, ip, userAgent, constants.LoginSuccessful)

	return &entities.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: rawRefreshToken,
		TokenType:    constants.BearerPrefix,
	}, nil
}

func (a *authService) RefreshToken(ctx context.Context, req *entities.RefreshTokenRequest, ip, userAgent string) (*entities.AuthResponse, error) {
	hash := utils.HashToken(req.RefreshToken)
	session, err := a.userSessionService.FindSessionByHash(ctx, hash)
	if err != nil || session.IsRevoked || session.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token tidak valid atau telah kedaluwarsa")
	}

	// Refresh Token Rotation: Revoke token lama
	_ = a.userSessionService.RevokeSessionByID(ctx, session.ID)

	user, err := a.authRepository.FindByID(ctx, session.UserID)
	if err != nil || !user.IsActive {
		return nil, errors.New("user tidak aktif atau tidak ditemukan")
	}

	// Buat Session & Refresh Token Baru
	newRawRefreshToken, _ := utils.GenerateRandomToken()
	newHash := utils.HashToken(newRawRefreshToken)

	newSession := &entities.UserSession{
		UserID:           user.ID,
		RefreshTokenHash: newHash,
		IpAddress:        ip,
		UserAgent:        userAgent,
		ExpiresAt:        time.Now().AddDate(0, 0, 7),
	}
	_ = a.userSessionService.CreateSession(ctx, newSession)

	jwtTTL, _ := strconv.Atoi(config.Config(constants.JwtTtlMinutes))
	if jwtTTL == 0 {
		jwtTTL = 15
	}

	newAccessToken, _ := utils.GenerateAccessToken(user.ID, user.Role, newSession.ID, config.Config(constants.EnvJwtSecret), jwtTTL)

	a.auditLogService.logAudit(ctx, &user.ID, constants.TokenRefreshed, ip, userAgent, constants.TokenRefreshedDetail)

	return &entities.AuthResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRawRefreshToken,
		TokenType:    constants.BearerPrefix,
	}, nil
}

func (a *authService) Logout(ctx context.Context, sessionID string) error {
	return a.userSessionService.RevokeSessionByID(ctx, sessionID)
}

func (a *authService) LogoutAll(ctx context.Context, userID string) error {
	return a.userSessionService.RevokeAllUserSessions(ctx, userID)
}

func (a *authService) GetMe(ctx context.Context, userID string) (*entities.User, error) {
	return a.authRepository.FindByID(ctx, userID)
}

func (a *authService) ForgotPassword(ctx context.Context, req *entities.ForgotPasswordRequest, ip, userAgent string) error {
	user, err := a.authRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil // Mencegah User Enumeration Attack
	}

	resetToken, _ := utils.GenerateRandomToken()
	expiresAt := time.Now().Add(30 * time.Minute)

	user.PasswordResetToken = utils.HashToken(resetToken)
	user.PasswordResetExpires = &expiresAt
	_ = a.authRepository.UpdateUser(ctx, user)

	a.auditLogService.logAudit(ctx, &user.ID, constants.ForgotPasswordRequested, ip, userAgent, constants.ResetPasswordDetail)

	// Panggil Email Service Anda di sini untuk mengirim resetToken ke email user
	return nil
}

func (a *authService) ResetPassword(ctx context.Context, req *entities.ResetPasswordRequest) error {
	hashedToken := utils.HashToken(req.Token)
	user, err := a.authRepository.FindByResetToken(ctx, hashedToken)
	if err != nil || user.PasswordResetExpires == nil || user.PasswordResetExpires.Before(time.Now()) {
		return errors.New("token reset password tidak valid atau kedaluwarsa")
	}

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.PasswordResetToken = ""
	user.PasswordResetExpires = nil
	_ = a.authRepository.UpdateUser(ctx, user)

	// Invalidasi semua session lama
	_ = a.userSessionService.RevokeAllUserSessions(ctx, user.ID)

	a.auditLogService.logAudit(ctx, &user.ID, constants.PasswordResetSuccess, "", "", constants.ChangePasswordDetail)
	return nil
}

func (a *authService) VerifyEmail(ctx context.Context, token string) error {
	hashedToken := utils.HashToken(token)
	user, err := a.authRepository.FindByVerificationToken(ctx, hashedToken)
	if err != nil {
		return errors.New("token verifikasi email tidak valid")
	}

	user.IsEmailVerified = true
	user.EmailVerificationToken = ""
	return a.authRepository.UpdateUser(ctx, user)
}
