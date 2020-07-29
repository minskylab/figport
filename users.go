package figport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type user struct {
	ID     string `json:"id"`
	Handle string `json:"handle"`
	Email  string `json:"email"`
}

func (fig *Figport) figmaUserProfileByToken(accessToken string) (*user, error) {
	endpoint, err := fig.figma.FigmaAPIURI("/v1/me")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	res, err := fig.httpClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	returnedUser := new(user)

	if err := json.NewDecoder(res.Body).Decode(returnedUser); err != nil {
		return nil, errors.WithStack(err)
	}

	if returnedUser.ID == "" {
		return nil, errors.New("invalid response from figma api")
	}

	logrus.WithFields(logrus.Fields{
		"id":     returnedUser.ID,
		"handle": returnedUser.Handle,
		"email":  returnedUser.Email,
	}).Debug("figma user profile response received")

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
