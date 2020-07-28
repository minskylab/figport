package figma

import (
	"os"

	"github.com/pkg/errors"
)

// ObtainImage obtein a node from your fileKey and generate a os.file ready to pipeline to another place
// Note that the strategy used in this algorithm is one request per one node, Fimga's API provide a render
// with multiple node but in this case I think is better make one by one because in this form we can inject
// processors and generate sub-renders in the future.
func (fig *Figma) ObtainImage(accessToken string, fileKey string, nodeID string, options RenderOptions) (*os.File, string, error) {
	render, err := fig.renderImageFromNode(accessToken, fileKey, []string{nodeID}, options)
	if err != nil {
		return nil, "", errors.WithStack(err)
	}

	if render.Err != "" {
		return nil, "", errors.New(render.Err)
	}

	if len(render.Images) == 0 {
		return nil, "", errors.New("invalid size of rendered images, Figma API bad response")
	}

	image := render.Images[nodeID]

	return fig.downloadFromFigmaRender(image)
}
