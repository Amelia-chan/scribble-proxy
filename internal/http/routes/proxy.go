package routes

import (
	"amelia-sh-proxy/internal/http"
	"github.com/gofiber/fiber/v2"
	"github.com/imroc/req/v3"
)

var RegularClient = req.C().SetBaseURL("https://scribblehub.com/")

var _ = http.Fiber.Attach(func(app *fiber.App) {
	app.Get("/*", func(c *fiber.Ctx) error {
		resp := RegularClient.Get(c.Params("*")).
			SetHeader("User-Agent", "Amelia/2.1 (Language=Kotlin/1.7.10, Developer=Shindou Mihou)").
			Do()
		if resp.Err != nil {
			return resp.Err
		}
		return c.Status(resp.StatusCode).
			Send(resp.Bytes())
	})
})
