package figport

import (
	"context"

	"github.com/minskylab/figport/exporting"
	"github.com/pkg/errors"
)

func (fig *Figport) processNodeName(nodeName string) exporting.ExportNodeOptions {
	return exporting.ExportNodeOptions{}
}

func (fig *Figport) executeDeployment(ctx context.Context, accessToken string, fileKey string) error {
	figmaFile, err := fig.figma.GetCompleteFile(accessToken, fileKey)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, artboard := range figmaFile.Document.Children {
		exportingOptions := fig.processNodeName(artboard.Name)

		for _, mod := range fig.mods {
			active := false
			for _, activeMod := range exportingOptions.Mods {
				if activeMod == mod.Name() {
					active = true
					break
				}
			}

			if !active {
				break
			}

			rendersOptions, err := mod.Process(exportingOptions, map[string]string{})
			if err != nil {
				return errors.WithStack(err)
			}

			for _, renderOptions := range rendersOptions {
				file, contentType, err := fig.figma.ObtainImage(accessToken, fileKey, artboard.ID, renderOptions)
				if err != nil {
					return errors.WithStack(err)
				}

				if _, err := fig.saveAsset(ctx, exportingOptions.Path, contentType, file); err != nil {
					return errors.WithStack(err)
				}
			}

		}

	}

	return nil
}
