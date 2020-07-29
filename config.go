package figport

import (
	"github.com/joho/godotenv"
	"github.com/minskylab/figport/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (fig *Figport) bootstrapDefaultConfig() error {
	godotenv.Load()

	fig.config.SetConfigName("figport.config")
	fig.config.SetConfigType("yaml")
	fig.config.AddConfigPath("/etc/figport/")
	fig.config.AddConfigPath(".")

	fig.config.SetDefault(config.PortKey, "8080")
	fig.config.SetDefault(config.HostNameKey, "127.0.0.1")
	fig.config.SetDefault(config.FigmaAPIBaseURL, "https://api.figma.com")
	fig.config.SetDefault(config.FigmaOauthURL, "https://www.figma.com")

	fig.config.SetEnvPrefix("FIGPORT")
	fig.config.AutomaticEnv()

	if err := fig.config.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}

	debugMode := fig.config.GetBool(config.DebugKey)

	if debugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.WithField("debug", debugMode).Info("configuration bootstrap done")

	return nil
}
