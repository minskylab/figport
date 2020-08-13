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
)

// New returns a new instance of a Figport controller
func New(ctx context.Context) (*Figport, error) {
	conf := viper.New()
	httpClient := &http.Client{
		Timeout: 15 * time.Second,
	}

	var database *Database
	var err error
	redisAddress := conf.GetString(config.RedisAddress)
	if redisAddress != "" {
		database, err = newDatabase(ctx, &redis.Options{
			Addr:     redisAddress,
			Password: "",
			DB:       0,
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	figmaHandler := figma.New(conf, httpClient)

	defaultMods := []Mod{
		&mods.SVGMod{},
		&mods.PNGMod{},
		&mods.JPGMod{},
		&mods.PDFMod{},
	}

	fiberApp := fiber.New()

	return &Figport{
		config:     conf,
		db:         database,
		figma:      figmaHandler,
		httpClient: httpClient,
		mods:       defaultMods,
		server:     fiberApp,
	}, nil
}
