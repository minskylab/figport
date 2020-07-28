package figport

import (
	"github.com/gofiber/fiber"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const webhooksRedisPrefix = "webhooks"

func (fig *Figport) registerWebhooks() {
	fig.server.Post("/webhooks/:id", func(c *fiber.Ctx) {
		webhookID := c.Params(":id", "")
		if webhookID == "" {
			sendError(c, errors.New("invalid webhook id"))
			return
		}

		// *
		secretKey := webhooksRedisPrefix + "." + webhookID + ".secret"
		secret, err := fig.db.redisClient.Get(c.Context(), secretKey).Result()
		if err != nil {
			sendError(c, errors.WithStack(err))
			return
		}

		querySecret := c.Query("secret", "")
		if secret != querySecret {
			sendError(c, errors.New("operation fordibben, your secret is not correct"))
			return
		}

		fileIDKey := webhooksRedisPrefix + "." + webhookID + ".filekey"
		fileKey, err := fig.db.redisClient.Get(c.Context(), fileIDKey).Result()
		if err != nil {
			sendError(c, errors.WithStack(err))
			return
		}

		userIDKey := webhooksRedisPrefix + "." + webhookID + ".owner"
		userID, err := fig.db.redisClient.Get(c.Context(), userIDKey).Result()
		if err != nil {
			sendError(c, errors.WithStack(err))
			return
		}

		tokenKey := userID + ".token.accesstoken"
		accessToken, err := fig.db.redisClient.Get(c.Context(), tokenKey).Result()
		if err != nil {
			sendError(c, errors.WithStack(err))
			return
		}

		if err := fig.executeDeployment(c.Context(), accessToken, fileKey); err != nil {
			sendError(c, errors.WithStack(err))
			return
		}

		logrus.Info("deployment executed")
	})

}
