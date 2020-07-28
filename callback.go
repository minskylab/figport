package figport

import (
	"context"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (fig *Figport) generateState(ctx context.Context) (string, error) {
	state := uuid.NewV4().String()
	value := time.Now().UTC().Format(time.RFC3339)
	res, err := fig.db.redisClient.Set(ctx, state, value, 30*time.Second).Result()
	if err != nil {
		return "", errors.WithStack(err)
	}
	return res, nil
}

func (fig *Figport) validateState(ctx context.Context, state string) error {
	stateIssuedTime, err := fig.db.redisClient.Get(ctx, state).Result()
	if err != nil {
		return errors.WithStack(err)
	}

	if stateIssuedTime == "" {
		return errors.New("invalid state, operation not allowed")
	}

	issuedTime, err := time.Parse(time.RFC3339, stateIssuedTime)
	if err != nil {
		return errors.WithStack(err)
	}

	if time.Now().Sub(issuedTime).Seconds() > 15 {
		return errors.New("your state value expire")
	}

	return nil
}

func (fig *Figport) destroyState(ctx context.Context, state string) error {
	// TODO: Handle the first parameter of result
	if _, err := fig.db.redisClient.Del(ctx, state).Result(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (fig *Figport) callback(ctx context.Context, code, state string) (string, error) {
	// TODO: In the future we need validate the code value too (better)
	if code != "" {
		return "", errors.New("invalid code responde, contact with figma")
	}

	if err := fig.validateState(ctx, state); err != nil {
		return "", errors.WithStack(err)
	}

	return code, nil
}
