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
	"github.com/valyala/fastjson"
)

func (fig *Figma) getFromFigmaFile(accessToken string, fileKey string, nodes ...string) (*File, error) {
	endpoint, err := fig.figmaURI("/v1/files", fileKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Add("X-FIGMA-TOKEN", accessToken)

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

// TODO: Fix the hardcoded names (e.g. X-FIGMA-TOKEN)
func (fig *Figma) renderImageFromNode(accessToken string, fileKey string, nodes []string, options RenderOptions) (*Render, error) {
	endpoint, err := fig.figmaURI("/v1/images", fileKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	uri, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := uri.Query()

	query.Add("ids", strings.Join(nodes, ","))
	query.Add("format", string(options.Format))
	query.Add("svg_include_id", strconv.FormatBool(options.SVGIncludeID))
	query.Add("svg_simplify_stroke", strconv.FormatBool(options.SVGSimplifyStroke))
	query.Add("use_absolute_bounds", strconv.FormatBool(options.UseAbsoluteBounds))
	query.Add("version", options.Version)

	uri.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodPost, uri.String(), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Add("X-FIGMA-TOKEN", accessToken)

	res, err := fig.httpClient.Do(req)
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

	image := map[string]string{}
	values.GetObject("images").Visit(func(key []byte, v *fastjson.Value) {
		if v != nil {
			image[string(key)] = v.String()
		}
	})

	figmaError := values.GetStringBytes("err")

	status := values.GetFloat64("status")

	return &Render{
		Err:    string(figmaError),
		Images: image,
		Status: status,
	}, nil
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
