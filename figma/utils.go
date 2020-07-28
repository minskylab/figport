package figma

import (
	"net/url"
	"path"

	"github.com/minskylab/figport/config"
	"github.com/pkg/errors"
)

func (fig *Figma) figmaURI(paths ...string) (string, error) {
	base := fig.config.GetString(config.FigmaOauthURL)
	u, err := url.Parse(base)
	if err != nil {
		return "", errors.WithStack(err)
	}

	u.Path = path.Join(append([]string{u.Path}, paths...)...)
	return u.String(), nil
}

// FigmaURI only esports figmaURI utility
func (fig *Figma) FigmaURI(paths ...string) (string, error) {
	return fig.figmaURI(paths...)
}
