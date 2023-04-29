package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const apiKeyHeader = "x-api-key"

func APIKeyAuth(apiKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientAPIKey := c.Get(apiKeyHeader)
		if clientAPIKey == "" {
			c.Status(http.StatusUnauthorized)
			return fmt.Errorf("missing API key")
		}

		if clientAPIKey != apiKey {
			c.Status(http.StatusForbidden)
			return fmt.Errorf("invalid API key")
		}

		return c.Next()
	}
}
