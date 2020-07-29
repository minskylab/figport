package figport

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"github.com/minskylab/figport/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type tokenResult struct {
	AccessToken  string `json:"access_token"`
	Expiration   int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func (fig *Figport) requestToken(code string) (*tokenResult, error) {
	base := fig.config.GetString(config.FigmaOauthURL)
	u, err := url.Parse(base)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	u.Path = path.Join(u.Path, "/api/oauth/token")

	q := u.Query()

	q.Add("client_id", fig.config.GetString(config.FigmaAppClientID))
	q.Add("redirect_uri", fig.config.GetString(config.FigmaRedirectURI))
	q.Add("client_secret", fig.config.GetString(config.FigmaClientSecret))
	q.Add("code", code)
	q.Add("grant_type", "authorization_code")

	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res, err := fig.httpClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tResult := new(tokenResult)
	if err := json.NewDecoder(res.Body).Decode(tResult); err != nil {
		return nil, errors.WithStack(err)
	}

	contentType := res.Header.Get("Content-Type")

	logrus.WithFields(logrus.Fields{
		"statusCode":  res.StatusCode,
		"contentType": contentType,
	}).Debug("figma response from /api/oauth/token")

	return tResult, nil
}
