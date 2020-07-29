package figport

import (
	"github.com/minskylab/figport/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (fig *Figport) bootstrapDefaultConfig() error {
	fig.config.SetConfigName("figport.config")
	fig.config.SetConfigType("yaml")
	fig.config.AddConfigPath("/etc/figport/")
	fig.config.AddConfigPath(".")

	fig.config.SetDefault(config.PortKey, "8080")
	fig.config.SetDefault(config.HostNameKey, "127.0.0.1")

	fig.config.SetEnvPrefix("figport")
	fig.config.AutomaticEnv()

	debugMode := fig.config.GetBool(config.DebugKey)
	logrus.Info(debugMode)
	if debugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if err := fig.config.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
