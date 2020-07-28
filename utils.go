package figport

import (
	"net/url"
	"path"

	"github.com/pkg/errors"
)

func (fig *Figport) figmaURI(paths ...string) (string, error) {
	base := fig.config.GetString(figmaOauthURL)
	u, err := url.Parse(base)
	if err != nil {
		return "", errors.WithStack(err)
	}

	u.Path = path.Join(append([]string{u.Path}, paths...)...)
	return u.String(), nil
}

// func (fig *Figport) withQuery(uri string, prevError error, ab string) (string, error) {
// 	if prevError != nil {
// 		return "", errors.WithStack(prevError)
// 	}

// 	return "", nil
// }
