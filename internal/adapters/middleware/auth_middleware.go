package middleware

import (
	"attendance-api/config"
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/enums"
	"attendance-api/internal/usecases"
	"attendance-api/internal/utils"
	"attendance-api/pkg/constants"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authRepo repositories.AuthRepository, userSessionService usecases.UserSessionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(constants.HeaderAuthorization)
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.BuildResponseUnauthorized("Header otorisasi diperlukan"))
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != constants.BearerPrefix {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.BuildResponseUnauthorized("Format token harus 'Bearer <token>'"))
			return
		}

		// Ini ambil ke ENV
		secret := config.Config(constants.EnvJwtSecret)
		claims, err := utils.ValidateToken(tokenParts[1], secret)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.BuildResponseUnauthorized("Token expired atau tidak valid"))
			return
		}
		log.Println("claims session id: ", claims.SessionID)

		sessionId := claims.SessionID

		userSession, err := userSessionService.FindSessionByID(c.Request.Context(), sessionId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildResponseFailed("User session id not found"))
			return
		}

		// Cek apakah session nya sudah logout atau belum
		if userSession.IsRevoked {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.BuildResponseFailed("Session has been logged out"))
			return
		}

		// simpan context untuk handler selanjutnya
		c.Set(constants.UserID, claims.UserID)
		c.Set(constants.UserRole, claims.Role)
		c.Set(constants.SessionID, claims.SessionID)

		c.Next()
	}
}

func RequireRoles(roles ...enums.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ini mengambil context ketika pengecekan lewat middleware
		roleVal, exists := c.Get(constants.UserRole)
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.BuildResponseForbidden("Role user tidak ditemukan"))
			return
		}

		userRole := roleVal.(enums.Role)
		log.Println("user role : ", userRole)
		hasRole := false
		for _, r := range roles {
			if r == userRole {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.BuildResponseForbidden("Anda tidak memiliki akses ke resource ini"))
			return
		}

		c.Next()
	}
}
