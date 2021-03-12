package figport

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/minskylab/figport/config"
	"github.com/sirupsen/logrus"
)

func (fig *Figport) bootstrapDefaultConfig(debug bool) error {
	_ = godotenv.Load()

	fig.config.SetConfigName("figport.config")
	fig.config.SetConfigType("yaml")
	fig.config.AddConfigPath("/etc/figport/")
	fig.config.AddConfigPath(".")

	// FigmaAPIBaseURL
	fig.config.SetDefault(config.PortKey, "8080")
	fig.config.SetDefault(config.HostNameKey, "127.0.0.1")
	fig.config.SetDefault(config.FigmaAPIBaseURL, "https://api.figma.com")
	fig.config.SetDefault(config.FigmaOauthURL, "https://www.figma.com")
	fig.config.SetDefault(config.FigportPrefix, "figport")

	// fig.config.SetDefault(config.S3Endpoint, "https://s3.amazonaws.com")
	// fig.config.SetDefault(config.S3Bucket, "figport")
	// fig.config.SetDefault(config.S3UseSSL, true)

	// deactivated because for this stage only need a stateless service.
	// fig.config.SetDefault(config.RedisAddress, "localhost:6379")

	fig.config.SetEnvPrefix("figport")
	fig.config.AutomaticEnv()
	fig.config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = fig.config.ReadInConfig() // try to read the config file (if exists)

	debugMode := fig.config.GetBool(config.DebugKey) || debug

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
		logrus.Warnf("global secret: \"%s\"", secret)
		logrus.Warn("that's so dangerous, try to set your own global secret by env variables or yaml config")
		fig.config.Set(config.GlobalSecret, secret)
	}

	logrus.WithFields(logrus.Fields{
		"debug":         debugMode,
		"personalToken": fig.withToken,
	}).Info("configuration bootstrapping done")

	return nil
}
