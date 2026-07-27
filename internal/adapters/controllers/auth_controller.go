package controllers

import (
	"attendance-api/internal/entities"
	"attendance-api/internal/usecases"
	"attendance-api/internal/utils"
	"attendance-api/pkg/constants"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	AuthController interface {
		CreateUser(ctx *gin.Context)
		Login(ctx *gin.Context)
		RefreshToken(ctx *gin.Context)
		Logout(ctx *gin.Context)
		LogoutAll(ctx *gin.Context)
		GetMe(ctx *gin.Context)
		ForgotPassword(ctx *gin.Context)
		ResetPassword(ctx *gin.Context)
		VerifyEmail(ctx *gin.Context)
	}

	authController struct {
		authService usecases.AuthService
	}
)

func NewAuthController(authService usecases.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (a *authController) CreateUser(ctx *gin.Context) {
	req := entities.CreateUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := a.authService.CreateUser(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("User berhasil dibuat", res))
}
func (a *authController) Login(ctx *gin.Context) {
	req := entities.LoginRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := a.authService.Login(ctx.Request.Context(), &req, ctx.ClientIP(), ctx.Request.UserAgent())
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Login Berhasil", res))
}

func (a *authController) RefreshToken(ctx *gin.Context) {
	req := entities.RefreshTokenRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	res, err := a.authService.RefreshToken(ctx.Request.Context(), &req, ctx.ClientIP(), ctx.Request.UserAgent())
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Token Berhasil Diperbarui", res))
}

func (a *authController) Logout(ctx *gin.Context) {
	sessionID := ctx.GetString("sessionID")

	if err := a.authService.Logout(ctx.Request.Context(), sessionID); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed("Gagal logout"))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Berhasil logout", nil))
}

func (a *authController) LogoutAll(ctx *gin.Context) {
	userID := ctx.GetString(constants.UserID)
	log.Println("user id : ", userID)
	if err := a.authService.LogoutAll(ctx.Request.Context(), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.BuildResponseFailed("Gagal logout seluruh perangkat"))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Berhasil logout dari seluruh perangkat", nil))
}

func (a *authController) GetMe(ctx *gin.Context) {
	// ini berasal dari context yg tersimpan di server
	// setelah melewati auth middleware tadi
	userID := ctx.GetString(constants.UserID)

	res, err := a.authService.GetMe(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.BuildResponseFailed("User tidak ditemukan"))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Success", res))
}

func (a *authController) ForgotPassword(ctx *gin.Context) {
	req := entities.ForgotPasswordRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	_ = a.authService.ForgotPassword(ctx.Request.Context(), &req, ctx.ClientIP(), ctx.Request.UserAgent())
	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Jika email terdaftar, instruksi reset password telah dikirim", nil))
}

func (a *authController) ResetPassword(ctx *gin.Context) {
	req := entities.ResetPasswordRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	if err := a.authService.ResetPassword(ctx.Request.Context(), &req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Password berhasil diperbarui. Silakan login kembali", nil))
}

func (a *authController) VerifyEmail(ctx *gin.Context) {
	token := ctx.Query(constants.Token)
	if token == "" {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed("Token tidak ditemukan"))
		return
	}

	if err := a.authService.VerifyEmail(ctx.Request.Context(), token); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Email berhasil diverifikasi", nil))
}
