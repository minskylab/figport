package figport

import (
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/minskylab/figport/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func sendError(c *fiber.Ctx, err error) error {
	logrus.Errorf("%#v", err)
	return c.JSON(map[string]string{
		"error": err.Error(),
	})
}

func (fig *Figport) registerAuth() {
	fig.server.Get("/auth", func(c *fiber.Ctx) error {
		state, err := fig.generateState(c.Context())
		if err != nil {
			return sendError(c, errors.WithStack(err))
		}

		figmaAuthURL, err := url.Parse(figmaBaseOAuthURL)
		if err != nil {
			return sendError(c, errors.WithStack(err))
		}

		q := figmaAuthURL.Query()

		q.Add("client_id", fig.config.GetString(config.FigmaAppClientID))
		q.Add("redirect_uri", fig.config.GetString(config.FigmaRedirectURI))
		q.Add("scope", "file_read")
		q.Add("state", state)
		q.Add("response_type", "code")

		figmaAuthURL.RawQuery = q.Encode()

		c.Redirect(figmaAuthURL.String(), http.StatusTemporaryRedirect)

		return nil
	})

	fig.server.Get("/oauth/callback", func(c *fiber.Ctx) error {
		body := c.Body()
		logrus.Info(body)

		code := c.Query("code", "00000000")
		state := c.Query("state", "")

		logrus.WithFields(logrus.Fields{
			"code":  code[:5] + "***",
			"state": state,
		}).Debug("oauth callback triggered")

		// fig.config.GetString(hostNameKey)

		code, err := fig.callback(c.Context(), code, state)
		if err != nil {
			return sendError(c, errors.WithStack(err))
		}

		token, err := fig.requestToken(code)
		if err != nil {
			return sendError(c, errors.WithStack(err))
		}

		user, err := fig.registerNewUser(c.Context(), token)
		if err != nil {
			return sendError(c, errors.WithStack(err))
		}

		// TODO: Implement a beauty user page response [200]

		c.SendString("Welcome " + user.Email)

		if err = fig.destroyState(c.Context(), state); err != nil {
			return sendError(c, errors.WithStack(err))
		}

		return nil
	})
}
