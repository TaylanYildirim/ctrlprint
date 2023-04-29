package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
)

// RateLimiter is a middleware to limit the rate of incoming requests.
func RateLimiter(requestsPerSecond int, burst int) fiber.Handler {
	clients := make(map[string]*rate.Limiter)
	mutex := sync.Mutex{}

	return func(c *fiber.Ctx) error {
		clientIP := c.IP()

		mutex.Lock()
		defer mutex.Unlock()

		limiter, exists := clients[clientIP]
		if !exists {
			limiter = rate.NewLimiter(rate.Every(time.Second/time.Duration(requestsPerSecond)), burst)
			clients[clientIP] = limiter
		}

		if !limiter.Allow() {
			c.Status(http.StatusTooManyRequests)
			return fmt.Errorf("too many requests")
		}

		return c.Next()
	}
}
