package responses

import (
	"amelia-sh-proxy/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger is a utility method that allows us to get a Logger that contains metadata about the
// request which  allows us to gain some more insights over the request that happened.
func Logger(ctx *fiber.Ctx) *zerolog.Logger {
	return utils.Ptr(log.With().
		Str("request", ctx.String()).
		Logger())
}
