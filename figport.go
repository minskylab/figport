package figport

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/minio/minio-go/v7"
	"github.com/minskylab/figport/figma"

	"github.com/spf13/viper"
	"github.com/valyala/fastjson"
)

// Figport is a struct to wrap all dependencies of Figport
type Figport struct {
	jsonParser *fastjson.Parser
	httpClient *http.Client
	config     *viper.Viper
	server     *fiber.App
	db         *Database
	s3Client   *minio.Client

	figma figma.Figma
	mods  []Mod
}
