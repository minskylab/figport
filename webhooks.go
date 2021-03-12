package figport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minskylab/figport/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// const webhooksRedisPrefix = "webhooks"

func (fig *Figport) registerDeploy() {
	fig.server.Post("/deploy/:fileKey", func(c *fiber.Ctx) error {
		if !fig.withToken {
			return sendError(c, errors.New("invalid configuration, you need set a personal token (alpha)"))
		}

		secret := c.Query("secret", "")
		if len(secret) != config.DefaultSecretSize || secret != fig.config.GetString(config.GlobalSecret) {
			return sendError(c, errors.New("invalid secret or not found, pass your secret correctly"))
		}

		fileKey := c.Params("fileKey")
		accessToken := fig.config.GetString(config.FigmaToken)

		logrus.WithFields(logrus.Fields{
			"fileKey": fileKey,
		}).Info("deploying figma file")

		var totalReports []nodeDeploymentReport

		report := make(chan nodeDeploymentReport, 10)

		if err := fig.executeDeployment(c.Context(), accessToken, fileKey, report); err != nil {
			return sendError(c, errors.WithStack(err))
		}

		for r := range report {
			totalReports = append(totalReports, r)
		}

		logrus.Info("deployment process done")

		return c.JSON(totalReports)
	})

}
