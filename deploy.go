package figport

import (
	"context"
	"regexp"
	"strconv"
	"strings"
	"time"

	"sync"

	"github.com/minskylab/figport/config"
	"github.com/minskylab/figport/exporting"
	"github.com/minskylab/figport/figma"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type nodeDeploymentReport struct {
	DeployAt            time.Time `json:"deployAt"`
	TotalRenders        int64 `json:"totalRenders"`
	TotalRendersSuccess int64 `json:"totalRendersSuccess"`
	TotalRendersFailed  int64    `json:"totalRendersFailed"`
	Renders             []string `json:"renders"` // TODO: Improve this Renders with a better struct (out of alpha)
}

func (fig *Figport) processNodeName(nodeName string) exporting.ExportNodeOptions {
	var activeMods []string
	var activeScales []float64
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

	logrus.WithFields(logrus.Fields{
		"params": params,
	}).Debug("extracting params")

	if params == "" {
		return map[string]string{}, nil
	}

	cleanedParams := strings.Trim(params, "({[]})")
	pairs := strings.Split(cleanedParams, ",")

	finalParams := map[string]string{}
	for _, pair := range pairs {
		nameValue := strings.Split(pair, "=")
		if len(nameValue) != 2 {
			return nil, errors.New("invalid param, please specify the name and the value in the form 'name=value'")
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

type asyncDeploymentParams struct {
	ctx context.Context
	rendersOptions []figma.RenderOptions
	exportingOptions exporting.ExportNodeOptions
	accessToken string
	fileKey string
	nodeID string
	prefix string
}

func (fig *Figport) asyncProcessAndDeployNodeRoutine(params asyncDeploymentParams, reportChannel chan<- nodeDeploymentReport) {
	group := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	totalNodes := int64(0)
	totalSuccess := int64(0)
	totalErrors := int64(0)

	var renders []string
	for _, renderOptions := range params.rendersOptions {
		group.Add(1)
		totalNodes += 1
		go func(wg *sync.WaitGroup, renderOptions figma.RenderOptions) {
			defer wg.Done()
			logrus.WithFields(logrus.Fields{
				"format":  renderOptions.Format,
				"scale":   renderOptions.Scale,
				"version": renderOptions.Version,
			}).Debug("rendering node with options")

			file, contentType, err := fig.figma.ObtainImage(params.accessToken, params.fileKey, params.nodeID, renderOptions)
			if err != nil {
				logrus.Error(errors.WithStack(err))
				mutex.Lock()
				totalErrors += 1
				mutex.Unlock()
				return
			}

			logrus.WithField("temp_file", file.Name()).Debug("temp_file generated from figma API render")

			cleanedPath := strings.TrimPrefix(params.exportingOptions.Path, params.prefix)
			cleanedPath = strings.TrimLeft(cleanedPath, " \\/.,")

			if renderOptions.Scale != 1.0 {
				cleanedPath += "@" + strconv.FormatFloat(renderOptions.Scale, 'f', -1, 64)
			}

			cleanedPath += "." + string(renderOptions.Format)

			if _, err := fig.saveAsset(params.ctx, cleanedPath, contentType, file); err != nil {
				logrus.Error(errors.WithStack(err))
				mutex.Lock()
				totalErrors += 1
				mutex.Unlock()
				return
			}
			mutex.Lock()
			renders = append(renders, cleanedPath)
			totalSuccess +=1
			mutex.Unlock()
		}(group, renderOptions)
	}
	group.Wait()
	report := nodeDeploymentReport{
		DeployAt:            time.Now(),
		TotalRenders:        totalNodes,
		TotalRendersSuccess: totalSuccess,
		TotalRendersFailed:  totalErrors,
		Renders:             renders,
	}
	reportChannel <- report
}

func (fig *Figport) deployNode(ctx context.Context, accessToken, fileKey string, nodeID, nodeName string) (<-chan nodeDeploymentReport, error) {
	exportingOptions := fig.processNodeName(nodeName)

	prefix := fig.config.GetString(config.FigportPrefix)
	if prefix == "" {
		return nil, errors.New("invalid prefix for enable your exports")
	}

	// To always add minimum @1 scale
	if len(exportingOptions.Scales) == 0 {
		exportingOptions.Scales = []float64{1.0}
	}

	// If we have more scale (e.g. [2.0]), that's if for always
	// have the basic scale (e.g. [1.0, 2.0]).
	if len(exportingOptions.Scales) > 0 {
		hasUnitScale := false
		for _, scale := range exportingOptions.Scales {
			if scale == 1.0 {
				hasUnitScale = true
				break
			}
		}
		if !hasUnitScale {
			exportingOptions.Scales = append(exportingOptions.Scales, 1.0)
		}
	}

	logrus.WithFields(logrus.Fields{
		"filename": exportingOptions.Filename,
		"mods":     exportingOptions.Mods,
		"scales":   exportingOptions.Scales,
		"path":     exportingOptions.Path,
		"raw":      exportingOptions.Raw,
	}).Debug("processing export naming")

	doneChannel := make(chan nodeDeploymentReport, 1)

	for _, activeMod := range fig.mods {
		modDescriptor := fig.getModIfIsActive(exportingOptions.Mods, activeMod.Name())
		if modDescriptor == "" {
			continue
		}

		params, err := fig.extractParamsFromModName(modDescriptor)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		logrus.WithFields(logrus.Fields{
			"mod":    activeMod.Name(),
			"params": params,
		}).Debug("processing with mod")

		rendersOptions, err := activeMod.Process(exportingOptions, params)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		logrus.WithFields(logrus.Fields{
			"totalRenders": len(rendersOptions),
		}).Debug("ready to save the Renders")

		go fig.asyncProcessAndDeployNodeRoutine(asyncDeploymentParams{
			ctx:              ctx,
			rendersOptions:   rendersOptions,
			exportingOptions: exportingOptions,
			accessToken:      accessToken,
			fileKey:          fileKey,
			nodeID:           nodeID,
			prefix:           prefix,
		}, doneChannel)
	}

	return doneChannel, nil
}

func (fig *Figport) executeDeployment(ctx context.Context, accessToken string, fileKey string, reportPipe chan nodeDeploymentReport) error {
	figmaFile, err := fig.figma.GetCompleteFile(accessToken, fileKey)
	if err != nil {
		return errors.WithStack(err)
	}

	prefix := fig.config.GetString(config.FigportPrefix)
	if prefix == "" {
		return errors.New("invalid prefix for enable your exports")
	}

	group := &sync.WaitGroup{}
	for nodeID, componentInfo := range figmaFile.Components {
		group.Add(1)
		go func (node string, reportPipe chan nodeDeploymentReport, componentInfo figma.Component, wg *sync.WaitGroup) {
			defer wg.Done()
			name := strings.ReplaceAll(componentInfo.Name, " ", "")

			toExport := strings.HasPrefix(name, prefix)
			if !toExport {
				return
			}

			logrus.WithFields(logrus.Fields{
				"name": name,
			}).Debug("reading component")

			deploymentDone, err := fig.deployNode(ctx, accessToken, fileKey, node, name)
			if err != nil {
				logrus.Error(errors.WithStack(err))
				return
			}
			report := <- deploymentDone
			reportPipe <- report
		}(nodeID, reportPipe, componentInfo, group)
	}

	group.Wait()
	close(reportPipe)

	return nil
}
