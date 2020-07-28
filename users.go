package figport

import (
	"context"
	"io/ioutil"

	"github.com/pkg/errors"
)

type user struct {
	ID     string
	Handle string
	Email  string
}

func (fig *Figport) figmaUserProfileByToken(token string) (*user, error) {
	endpoint, err := fig.figmaURI("/v1/me")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res, err := fig.httpClient.Get(endpoint)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	values, err := fig.jsonParser.ParseBytes(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	returnedUser := new(user)

	returnedUser.ID = string(values.GetStringBytes("id"))
	returnedUser.Email = string(values.GetStringBytes("email"))
	returnedUser.Handle = string(values.GetStringBytes("handle"))

	return returnedUser, nil
}

func (fig *Figport) registerNewUser(ctx context.Context, token *tokenResult) (*user, error) {
	u, err := fig.figmaUserProfileByToken(token.AccessToken)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// saving token ->  userid
	if _, err := fig.db.redisClient.Set(ctx, token.AccessToken, u.ID, 0).Result(); err != nil {
		return nil, errors.WithStack(err)
	}

	prefix := u.ID

	//////// Profile Related

	// saving userid.id ->  id
	key := prefix + ".id"
	if _, err := fig.db.redisClient.Set(ctx, key, u.ID, 0).Result(); err != nil {
		return nil, errors.WithStack(err)
	}

	// saving userid.handle ->  email
	key = prefix + ".handle"
	if _, err := fig.db.redisClient.Set(ctx, key, u.Handle, 0).Result(); err != nil {
		return nil, errors.WithStack(err)
	}

	// saving userid.email ->  handle
	key = prefix + ".email"
	if _, err := fig.db.redisClient.Set(ctx, key, u.Email, 0).Result(); err != nil {
		return nil, errors.WithStack(err)
	}

	//////// Token related

	// saving userid.token.accesstoken ->  token.AccessToken
	key = prefix + ".token.accesstoken"
	if _, err := fig.db.redisClient.Set(ctx, key, token.AccessToken, 0).Result(); err != nil {
		return nil, errors.WithStack(err)
	}

	// saving userid.token.expiration ->  token.Expiration
	key = prefix + ".token.expiration"
	if _, err := fig.db.redisClient.Set(ctx, key, token.Expiration, 0).Result(); err != nil {
		return nil, errors.WithStack(err)
	}

	// saving userid.token.refreshtoken ->  token.RefreshToken
	key = prefix + ".token.refreshtoken"
	if _, err := fig.db.redisClient.Set(ctx, key, token.RefreshToken, 0).Result(); err != nil {
		return nil, errors.WithStack(err)
	}

	return u, nil
}
