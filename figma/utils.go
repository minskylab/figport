package figma

import (
	"net/url"
	"path"

	"github.com/minskylab/figport/config"
	"github.com/pkg/errors"
)

func (fig *Figma) figmaURI(base string, paths ...string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", errors.WithStack(err)
	}

	u.Path = path.Join(append([]string{u.Path}, paths...)...)
	return u.String(), nil
}

// FigmaURI only exports figmaURI utility
func (fig *Figma) FigmaURI(paths ...string) (string, error) {
	base := fig.config.GetString(config.FigmaOauthURL)
	return fig.figmaURI(base, paths...)
}

// FigmaAPIURI uses the api base instead to oauth base
func (fig *Figma) FigmaAPIURI(paths ...string) (string, error) {
	base := fig.config.GetString(config.FigmaAPIBaseURL)
	return fig.figmaURI(base, paths...)
}
