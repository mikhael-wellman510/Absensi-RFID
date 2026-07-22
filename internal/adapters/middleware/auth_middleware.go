package middleware

import (
	"attendance-api/config"
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/enums"
	"attendance-api/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authRepo repositories.AuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Header otorisasi diperlukan"})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Format token harus 'Bearer <token>'"})
			return
		}

		// Ini ambil ke ENV
		secret := config.Config("JWT_SECRET")
		claims, err := utils.ValidateToken(tokenParts[1], secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired atau tidak valid"})
			return
		}

		// simpan context untuk handler selanjutnya
		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Set("sessionID", claims.SessionID)

		c.Next()
	}
}

func RequireRoles(roles ...enums.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ini mengambil context ketika pengecekan lewat middleware
		roleVal, exists := c.Get("userRole")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Role user tidak ditemukan"})
			return
		}

		userRole := roleVal.(enums.Role)
		hasRole := false
		for _, r := range roles {
			if r == userRole {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Anda tidak memiliki akses ke resource ini"})
			return
		}

		c.Next()
	}
}
