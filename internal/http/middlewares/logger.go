package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func Log(c *fiber.Ctx) (err error) {
	start := time.Now()
	msg := c.Method() + " " + c.Path()
	err = c.Next()
	code := c.Response().StatusCode()

	builder := log.With().
		Int("status", code).
		Str("ip", c.IP()).
		Str("latency", time.Since(start).String()).
		Str("user-agent", c.Get(fiber.HeaderUserAgent)).
		Str("request-id", c.Locals("requestid").(string))

	if err != nil {
		builder = builder.Str("err", err.Error())
	}

	logger := builder.Logger()
	switch {
	case code >= fiber.StatusBadRequest && code < fiber.StatusInternalServerError:
		logger.Warn().Msg(msg)
	case code >= http.StatusInternalServerError:
		logger.Error().Msg(msg)
	default:
		logger.Info().Msg(msg)
	}
	return err
}
