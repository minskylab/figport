package figport

import (
	"errors"

	"github.com/gofiber/fiber"
)

func (fig *Figport) registerWebhooks() {
	fig.server.Post("/webhooks/:id", func(c *fiber.Ctx) {
		webhookID := c.Params(":id", "")
		if webhookID == "" {
			sendError(c, errors.New("invalid webhook id"))
			return
		}

	})

}
