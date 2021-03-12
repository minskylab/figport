package figport

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minskylab/figport/config"
	"github.com/minskylab/figport/structure"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (fig *Figport) registerStructure() {
	fig.server.Post("/structure/:fileKey", func(c *fiber.Ctx) error {
		if !fig.withToken {
			return sendError(c, errors.New("invalid configuration, you need set a personal token (alpha)"))
		}

		secret := c.Query("secret", "")

		if secret != fig.config.GetString(config.GlobalSecret) {
			return sendError(c, errors.New("invalid secret or not found, pass your secret correctly"))
		}

		fileKey := c.Params("fileKey")
		accessToken := fig.config.GetString(config.FigmaToken)

		logrus.WithFields(logrus.Fields{
			"fileKey": fileKey,
		}).Info("finding structure...")

		t0 := time.Now()

		figmaFile, err := fig.figma.GetCompleteFile(accessToken, fileKey)
		if err != nil {
			return errors.WithStack(err)
		}

		routes := structure.ExtractStructureFromFigmaFile(figmaFile)

		logrus.WithFields(logrus.Fields{
			"duration": time.Since(t0),
		}).Info("structure extraction done")

		return c.JSON(routes)
	})
}
