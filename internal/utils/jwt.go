package utils

import (
	"attendance-api/internal/enums"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID    string     `json:"userId"`
	Role      enums.Role `json:"role"`
	SessionID string     `json:"sessionId"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID string, role enums.Role, sessionID string, secret string, ttlMinutes int) (string, error) {
	claims := JWTClaims{
		UserID:    userID,
		Role:      role,
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ttlMinutes) * time.Minute)),
			// Kalau mau testing JWT 10 detik
			//ExpiresAt: jwt.NewNumericDate(
			//	time.Now().Add(10 * time.Second),
			//),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenStr string, secret string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("method signing token tidak valid")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return claims, nil
}
