package figma

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (fig *Figma) getFromFigmaFile(accessToken string, fileKey string) (*File, error) {
	endpoint, err := fig.FigmaAPIURI("/v1/files", fileKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(accessToken) > 40 { // Personal Access Token
		req.Header.Add("X-FIGMA-TOKEN", accessToken)
	} else { // OAuth generated Token
		req.Header.Add("Authorization", "Bearer "+accessToken)
	}

	logrus.WithFields(logrus.Fields{
		"endpoint": endpoint,
	}).Debug("getting figma file")

	res, err := fig.httpClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	file := new(File)
	if err := json.NewDecoder(res.Body).Decode(file); err != nil {
		return nil, errors.WithStack(err)
	}

	return file, nil
}

// TODO: Fix the hardcoded names
func (fig *Figma) renderImageFromNode(accessToken string, fileKey string, nodes []string, options RenderOptions) (*Render, error) {
	endpoint, err := fig.FigmaAPIURI("/v1/images", fileKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	uri, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := uri.Query()

	if options.Scale == 0.0 {
		options.Scale = 1.0
	}

	if options.Scale < 0.0 {
		options.Scale *= -1
	}

	query.Add("ids", strings.Join(nodes, ","))
	query.Add("format", string(options.Format))
	query.Add("scale", strconv.FormatFloat(options.Scale, 'f', -1, 64))
	query.Add("svg_include_id", strconv.FormatBool(options.SVGIncludeID))
	query.Add("svg_simplify_stroke", strconv.FormatBool(options.SVGSimplifyStroke))
	query.Add("use_absolute_bounds", strconv.FormatBool(options.UseAbsoluteBounds))
	query.Add("version", options.Version)

	uri.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(accessToken) > 40 { // Personal Access Token
		req.Header.Add("X-FIGMA-TOKEN", accessToken)
	} else { // OAuth generated Token
		req.Header.Add("Authorization", "Bearer "+accessToken)
	}

	res, err := fig.httpClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	render := new(Render)

	if err := json.NewDecoder(res.Body).Decode(render); err != nil {
		return nil, errors.WithStack(err)
	}

	return render, nil
}

func (fig *Figma) downloadFromFigmaRender(imageURL string) (*os.File, string, error) {
	res, err := fig.httpClient.Get(imageURL)
	if err != nil {
		return nil, "", errors.WithStack(err)
	}

	contentType := res.Header.Get("Content-Type")

	logrus.WithFields(logrus.Fields{
		"contentLength": res.ContentLength,
		"imageURL":      imageURL,
	}).Debug("image downloaded from figma render")

	file, err := ioutil.TempFile(os.TempDir(), "figport")
	if err != nil {
		return nil, "", errors.WithStack(err)
	}

	if _, err := io.Copy(file, res.Body); err != nil {
		return nil, "", errors.WithStack(err)
	}

	return file, contentType, nil
}
