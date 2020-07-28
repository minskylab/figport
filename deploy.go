package figport

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/minskylab/figport/config"
	"github.com/minskylab/figport/exporting"
	"github.com/pkg/errors"
)

func (fig *Figport) processNodeName(nodeName string) exporting.ExportNodeOptions {
	activeMods := []string{}
	activeScales := []float64{}
	path := nodeName

	nodeName = strings.TrimSpace(nodeName)

	mods := regexp.MustCompile(`:[\w()>,.= ]+`).FindAllString(nodeName, -1)
	for _, mod := range mods {
		activeMods = append(activeMods, strings.ReplaceAll(mod, ":", ""))
		path = strings.ReplaceAll(path, mod, "")
	}

	scales := regexp.MustCompile(`@[\w.]+`).FindAllString(nodeName, -1)
	for _, scale := range scales {
		floatScale, err := strconv.ParseFloat(strings.ReplaceAll(scale, "@", ""), 64)
		if err != nil {
			continue
		}
		activeScales = append(activeScales, floatScale)
		path = strings.ReplaceAll(path, scale, "")
	}

	parts := strings.Split(path, "/")
	if len(parts) < 1 {
		return exporting.ExportNodeOptions{}
	}

	filename := parts[len(parts)-1]

	return exporting.ExportNodeOptions{
		Mods:     activeMods,
		Scales:   activeScales,
		Filename: filename,
		Path:     path,
		Raw:      nodeName,
	}
}

func (fig *Figport) extractParamsFromModName(name string) (map[string]string, error) {
	extractParams := regexp.MustCompile(`\([\w >,.= ]+\)`)
	params := extractParams.FindString(name)
	cleanedParams := strings.Trim(params, "({[]})")
	pairs := strings.Split(cleanedParams, ",")

	finalParams := map[string]string{}
	for _, pair := range pairs {
		nameValue := strings.Split(pair, "=")
		if len(nameValue) != 2 {
			return nil, errors.New("invalid param, please especify the name and the value in the form 'name=value'")
		}
		finalParams[nameValue[0]] = nameValue[1]
	}

	return finalParams, nil
}

func (fig *Figport) getModIfIsActive(mods []string, activeMod string) string {
	for _, modDescription := range mods {
		if strings.HasPrefix(modDescription, activeMod) {
			return modDescription
		}
	}
	return ""
}

func (fig *Figport) deployNode(ctx context.Context, accessToken, fileKey string, nodeID, nodeName string) error {
	exportingOptions := fig.processNodeName(nodeName)

	for _, activeMod := range fig.mods {
		modDescriptor := fig.getModIfIsActive(exportingOptions.Mods, activeMod.Name())
		if modDescriptor == "" {
			continue
		}

		params, err := fig.extractParamsFromModName(modDescriptor)
		if err != nil {
			return errors.WithStack(err)
		}

		rendersOptions, err := activeMod.Process(exportingOptions, params)
		if err != nil {
			return errors.WithStack(err)
		}

		for _, renderOptions := range rendersOptions {
			file, contentType, err := fig.figma.ObtainImage(accessToken, fileKey, nodeID, renderOptions)
			if err != nil {
				return errors.WithStack(err)
			}

			if _, err := fig.saveAsset(ctx, exportingOptions.Path, contentType, file); err != nil {
				return errors.WithStack(err)
			}
		}

	}

	return nil
}

func (fig *Figport) executeDeployment(ctx context.Context, accessToken string, fileKey string) error {
	figmaFile, err := fig.figma.GetCompleteFile(accessToken, fileKey)
	if err != nil {
		return errors.WithStack(err)
	}

	prefix := fig.config.GetString(config.FigportPrefix)
	if prefix == "" {
		return errors.New("invalid prefix to enable your exportations")
	}

	for _, artboard := range figmaFile.Document.Children {
		name := strings.ReplaceAll(artboard.Name, " ", "")

		if !strings.HasPrefix(name, prefix) {
			continue
		}

		if err := fig.deployNode(ctx, accessToken, fileKey, artboard.ID, name); err != nil {
			return errors.WithStack(err)
		}
	}

	for nodeID, componentInfo := range figmaFile.Components {
		name := strings.ReplaceAll(componentInfo.Name, " ", "")

		if !strings.HasPrefix(strings.ReplaceAll(componentInfo.Name, " ", ""), prefix) {
			continue
		}

		if err := fig.deployNode(ctx, accessToken, fileKey, nodeID, name); err != nil {
			return errors.WithStack(err)
		}

	}

	return nil
}
