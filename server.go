package figport

import (
	"strings"

	"github.com/minskylab/figport/config"
	"github.com/sirupsen/logrus"
)

func (fig *Figport) runServer() error {
	fig.registerAuth()
	fig.registerDeploy()

	port := fig.config.GetString(config.PortKey)
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	logrus.WithFields(logrus.Fields{
		"port": port,
	}).Info("server ready to listen")

	fig.server.Settings.DisableStartupMessage = true

	return fig.server.Listen(port)
}
