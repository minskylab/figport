package figport

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (fig *Figport) registerK8SLiveness() {
	fig.server.Get("/live", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("systems ok")
	})

	fig.server.Get("/ready", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("systems ok")
	})
}
