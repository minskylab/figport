package figport

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/minio/minio-go/v7"
	"github.com/minskylab/figport/config"
	"github.com/minskylab/figport/figma"

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
	s3Client   *minio.Client

	figma figma.Figma
	mods  []Mod
}

func (fig *Figport) Start() error {
	port := fig.config.GetString(config.Port)
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return fig.server.Listen(port)
}
