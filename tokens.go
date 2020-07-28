package figport

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/minskylab/figport/config"
	"github.com/pkg/errors"
)

type tokenResult struct {
	AccessToken  string
	Expiration   string
	RefreshToken string
}

func (fig *Figport) requestToken(code string) (*tokenResult, error) {
	base := fig.config.GetString(config.FigmaOauthURL)
	u, err := url.Parse(base)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	u.Path = path.Join(u.Path, "/token")

	q := u.Query()

	q.Add("client_id", fig.config.GetString(config.FigmaOauthURL))
	q.Add("redirect_uri", fig.config.GetString(config.FigmaRedirectURI))
	q.Add("client_secret", fig.config.GetString(config.FigmaClientSecret))
	q.Add("code", code)
	q.Add("response_type", "authorization_code")

	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res, err := fig.httpClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	vals, err := fig.jsonParser.ParseBytes(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// TODO: Save refresh and expiration

	accessToken := vals.GetStringBytes("access_token")
	expiration := vals.GetStringBytes("expires_in")
	refreshToken := vals.GetStringBytes("refresh_token")

	return &tokenResult{
		AccessToken:  string(accessToken),
		Expiration:   string(expiration),
		RefreshToken: string(refreshToken),
	}, nil
}
