package middlewares

import (
	"amelia-sh-proxy/internal/responses"
	"amelia-sh-proxy/pkg/env"
	"github.com/gofiber/fiber/v2"
)

func RequireAuthentication(c *fiber.Ctx) (err error) {
	if c.Get("Authorization") != env.Secret.Get() {
		responses.UnknownClient.Reply(c)
		return nil
	}
	return c.Next()
}
