package figport

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minskylab/figport/figma"
	"github.com/minskylab/figport/mods"
	"github.com/spf13/viper"
)

// New returns a new instance of a Figport controller
func New(ctx context.Context) (*Figport, error) {
	conf := viper.New()
	httpClient := &http.Client{
		Timeout: 5 * 60 * time.Second,
	}

	var database *Database
	// var err error
	// redisAddress := conf.GetString(config.RedisAddress)
	// if redisAddress != "" {
	// 	database, err = newDatabase(ctx, &redis.Options{
	// 		Addr:     redisAddress,
	// 		Password: "",
	// 		DB:       0,
	// 	})
	// 	if err != nil {
	// 		return nil, errors.WithStack(err)
	// 	}
	// }

	figmaHandler := figma.New(conf, httpClient)

	defaultMods := []Mod{
		&mods.SVGMod{},
		&mods.PNGMod{},
		&mods.JPGMod{},
		&mods.PDFMod{},
	}

	fiberApp := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	return &Figport{
		config:     conf,
		db:         database,
		figma:      figmaHandler,
		httpClient: httpClient,
		mods:       defaultMods,
		server:     fiberApp,
	}, nil
}
