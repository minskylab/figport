package figport

import (
	"github.com/gofiber/fiber"
	"github.com/minskylab/figport/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// const webhooksRedisPrefix = "webhooks"

func (fig *Figport) registerDeploy() {
	fig.server.Post("/deploy/:fileKey", func(c *fiber.Ctx) {
		if !fig.withToken {
			sendError(c, errors.New("invalid configuration, you need set a personal token (alpha)"))
			return
		}

		secret := c.Query("secret", "")
		if len(secret) != config.DefaultSecretSize || secret != fig.config.GetString(config.GlobalSecret)  {
			sendError(c, errors.New("invalid secret or not found, pass your secret correctly"))
			return
		}

		fileKey := c.Params("fileKey")
		accessToken := fig.config.GetString(config.FigmaToken)

		logrus.WithFields(logrus.Fields{
			"fileKey": fileKey,
		}).Info("deploying figma file")

		if err := fig.executeDeployment(c.Context(), accessToken, fileKey); err != nil {
			sendError(c, errors.WithStack(err))
			return
		}

		logrus.Info("deployment executed")
	})

}
