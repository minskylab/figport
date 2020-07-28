package figma

import (
	"net/http"

	"github.com/spf13/viper"
	"github.com/valyala/fastjson"
)

// Figma wraps the incomplete figma client
type Figma struct {
	config     *viper.Viper
	httpClient *http.Client
	jsonParser *fastjson.Parser
}
