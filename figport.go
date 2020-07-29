package figport

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minskylab/figport/config"
	"github.com/minskylab/figport/figma"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	"github.com/valyala/fastjson"
)

// Figport is a struct to wrap all dependencies of Figport
type Figport struct {
	withToken  bool
	jsonParser *fastjson.Parser
	httpClient *http.Client
	config     *viper.Viper
	server     *fiber.App
	db         *Database
	s3Options  *s3storageOptions
	s3Client   *minio.Client

	figma figma.Figma
	mods  []Mod
}

// Start bootstraps the start actions
func (fig *Figport) Start() error {
	fig.config.SetConfigName("figport.config.yaml")
	fig.config.AddConfigPath("/etc/figport/")
	fig.config.AddConfigPath(".")

	debugMode := fig.config.GetBool("figport.debug")

	if debugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.WithField("debug", debugMode).Info("starting sigport service")

	if err := fig.config.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}

	fig.config.SetEnvPrefix("figport")
	fig.config.AutomaticEnv()

	// fig.config.SetDefault(, value interface{})

	port := fig.config.GetString(config.Port)
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	// s3 connection

	client, err := minio.New(fig.s3Options.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(fig.s3Options.AccessKeyID, fig.s3Options.SecretKey, ""),
		Region: fig.s3Options.Region,
		Secure: fig.s3Options.UseSSL,
	})

	if err != nil {
		logrus.Panic(errors.WithStack(err))
	}

	fig.s3Client = client

	// server start

	return fig.server.Listen(port)
}
