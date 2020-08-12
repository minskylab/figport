package figport

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/minskylab/figport/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (fig *Figport) bootstrapDefaultConfig() error {
	_ = godotenv.Load()

	fig.config.SetConfigName("figport.config")
	fig.config.SetConfigType("yaml")
	fig.config.AddConfigPath("/etc/figport/")
	fig.config.AddConfigPath(".")

	fig.config.SetDefault(config.PortKey, "8080")
	fig.config.SetDefault(config.HostNameKey, "127.0.0.1")
	fig.config.SetDefault(config.FigmaAPIBaseURL, "https://api.figma.com")
	fig.config.SetDefault(config.FigmaOauthURL, "https://www.figma.com")
	fig.config.SetDefault(config.FigportPrefix, "figport")

	fig.config.SetDefault(config.RedisAddress, "localhost:6379")

	fig.config.SetEnvPrefix("figport")
	fig.config.AutomaticEnv()
	fig.config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := fig.config.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}

	debugMode := fig.config.GetBool(config.DebugKey)

	if debugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}

	token := fig.config.GetString(config.FigmaToken)
	if token != "" {
		fig.withToken = true
	} else {
		fig.withToken = false
	}

	globalSecret := fig.config.GetString(config.GlobalSecret)
	if globalSecret == "" {
		logrus.Info("global secret not manually choose")
		logrus.Info("generating a new random global secret")
		secret := newRandomString(config.DefaultSecretSize)
		logrus.Infof("global secret: \"%s\"", secret)
		logrus.Warn("that's so dangerous, try to set your own global secret by env variables ot yaml config")
		fig.config.Set(config.GlobalSecret, secret)
	}

	logrus.WithFields(logrus.Fields{
		"debug":         debugMode,
		"personalToken": fig.withToken,
	}).Info("configuration bootstrap done")

	return nil
}
