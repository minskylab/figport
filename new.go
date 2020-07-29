package figport

import (
	"context"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber"
	"github.com/minskylab/figport/config"
	"github.com/minskylab/figport/figma"
	"github.com/minskylab/figport/mods"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/valyala/fastjson"
)

// NewDefault ...
func NewDefault(ctx context.Context, withToken bool) (*Figport, error) {
	viper := viper.New()
	httpClient := &http.Client{
		Timeout: 15 * time.Second,
	}

	address := "localhost:6379"
	if addr := viper.GetString(config.FigmaRedirectURI); addr != "" {
		address = addr
	}

	jsonParser := &fastjson.Parser{}

	database, err := newDatabase(ctx, &redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	figmaHandler := figma.New(viper, httpClient, jsonParser)

	defaultMods := []Mod{
		&mods.SVGMod{},
	}

	fiberApp := fiber.New()

	return &Figport{
		withToken:  withToken,
		jsonParser: jsonParser,
		config:     viper,
		db:         database,
		figma:      figmaHandler,
		httpClient: httpClient,
		mods:       defaultMods,
		server:     fiberApp,
	}, nil
}
