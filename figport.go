package figport

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/minio/minio-go/v7"
	"github.com/minskylab/figport/figma"
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
	// configuration
	if err := fig.bootstrapDefaultConfig(); err != nil {
		logrus.Panic(err)
	}

	// s3 connection
	if err := fig.connectS3(); err != nil {
		logrus.Panic(err)
	}

	// server start
	if err := fig.runServer(); err != nil {
		logrus.Panic(err)
	}

	return nil
}
