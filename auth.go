package figport

import (
	"net/http"
	"net/url"

	"github.com/gofiber/fiber"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func sendError(c *fiber.Ctx, err error) {
	logrus.Error(err.Error())
	c.JSON(map[string]string{
		"error": err.Error(),
	})
}

func (fig *Figport) registerAuth() {
	fig.server.Get("/auth", func(c *fiber.Ctx) {
		state, err := fig.generateState(c.Context())
		if err != nil {
			sendError(c, errors.WithStack(err))
			return
		}

		figmaAuthURL, err := url.Parse(figmaBaseOAuthURL)
		if err != nil {
			sendError(c, errors.WithStack(err))
			return
		}

		q := figmaAuthURL.Query()

		q.Add("client_id", fig.config.GetString(figmaOauthURL))
		q.Add("redirect_uri", fig.config.GetString(figmaRedirectURI))
		q.Add("scope", "file_read")
		q.Add("state", state)
		q.Add("response_type", "code")

		figmaAuthURL.RawQuery = q.Encode()

		c.Redirect(figmaAuthURL.String(), http.StatusTemporaryRedirect)

	})

	fig.server.Get("/oauth/callback", func(c *fiber.Ctx) {
		body := c.Body()
		logrus.Info(body)

		code := c.Query("code", "")
		state := c.Query("state", "")

		logrus.Info(code, state)

		// fig.config.GetString(hostNameKey)

		code, err := fig.callback(c.Context(), code, state)
		if err != nil {
			sendError(c, errors.WithStack(err))
			return
		}

		token, err := fig.requestToken(code)
		if err != nil {
			sendError(c, errors.WithStack(err))
			return
		}

		user, err := fig.registerNewUser(c.Context(), token)
		if err != nil {
			sendError(c, errors.WithStack(err))
		}

		// TODO: Implement a beauty user page response [200]

		c.SendString("Welcome " + user.Email)
	})
}
