package http

import (
	"amelia-sh-proxy/internal/http/middlewares"
	"amelia-sh-proxy/internal/responses"
	"amelia-sh-proxy/pkg/clients"
	"amelia-sh-proxy/pkg/env"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/dchest/uniuri"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog/log"
)

type iFiber struct {
	clients.Client
	App         *fiber.App
	Executables []Executable
}

type Executable = func(app *fiber.App)

var Fiber = iFiber{
	App: fiber.New(fiber.Config{
		Prefork:                 false,
		ServerHeader:            "Amelia-Proxy",
		AppName:                 "Amelia-chan",
		EnableTrustedProxyCheck: true,
		JSONDecoder:             sonic.Unmarshal,
		JSONEncoder:             sonic.Marshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			var e *fiber.Error
			if errors.As(err, &e) {
				return ctx.SendStatus(e.Code)
			}
			responses.HandleErr(ctx, err)
			return nil
		},
		Immutable: true,
	}),
	Executables: []Executable{},
}

func (client *iFiber) Attach(executable Executable) bool {
	client.Executables = append(client.Executables, executable)
	return true
}

func (client *iFiber) Init() {
	env.Secret.MustGet()
	logger := log.With().Str("client", "Fiber").Logger()
	client.App.Use(
		recover.New(recover.Config{EnableStackTrace: true}),
		requestid.New(requestid.Config{
			Generator: func() string {
				return uniuri.NewLen(64)
			},
			Header:     "X-Request-Id",
			ContextKey: "requestid",
			Next:       nil,
		}),
		middlewares.RequireAuthentication,
		middlewares.Log,
	)
	for _, executable := range client.Executables {
		executable(client.App)
	}
	if err := client.App.Listen(":7631"); err != nil {
		logger.Panic().Err(err).Msg("Failed to start server")
		return
	}
}
