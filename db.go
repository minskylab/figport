package figport

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Database ...
type Database struct {
	redisClient *redis.Client
}

func newDatabase(ctx context.Context, opts *redis.Options) (*Database, error) {
	rdb := redis.NewClient(opts)
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, errors.WithStack(err)
	}

	logrus.WithField("address", opts.Addr).Debug("redis connection entableshid")

	return &Database{redisClient: rdb}, nil
}
