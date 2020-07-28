package figport

import (
	"github.com/gofiber/fiber"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const webhooksRedisPrefix = "webhooks"

func (fig *Figport) registerWebhooks() {
	fig.server.Post("/webhooks/:id", func(c *fiber.Ctx) {
		webhookID := c.Params(":id", "")
		if webhookID == "" {
			sendError(c, errors.New("invalid webhook id"))
			return
		}

		key := webhooksRedisPrefix + "." + webhookID

		secretHash, err := fig.db.redisClient.Get(c.Context(), key).Result()
		if err != nil {
			sendError(c, errors.WithStack(err))
			return
		}

		querySecret := c.Query("secret", "")

		if err := bcrypt.CompareHashAndPassword([]byte(secretHash), []byte(querySecret)); err != nil {
			sendError(c, errors.WithStack(err))
			return
		}

	})

}
