package middleware

import (
	"attendance-api/internal/utils"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiters = make(map[string]*rate.Limiter)
var mu sync.Mutex

func getVisitorLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := limiters[ip]
	if !exists {
		// Izinkan maks 5 request per detik, burst hingga 10 request
		limiter = rate.NewLimiter(5, 10)
		limiters[ip] = limiter
	}

	return limiter
}

// RateLimiter Ini untuk mencegah brute force
func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := getVisitorLimiter(ip)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, utils.BuildResponseTooManyRequests("Maksimum rate limit terlampaui. Silakan coba beberapa saat lagi."))
			return
		}

		c.Next()
	}
}
