package mods

import (
	"strconv"

	"github.com/minskylab/figport/exporting"
	"github.com/minskylab/figport/figma"
)

type SVGMod struct{}

func (mod *SVGMod) Name() string {
	return "svg"
}

func (mod *SVGMod) Params() map[string]string {
	return map[string]string{
		"includeID":      "[boolean] if it's true, your svg file will have the figma node id",
		"simplifyStroke": "[boolean] internal Figma stroke simplifier activation",
	}
}

func (mod *SVGMod) Process(opts exporting.ExportNodeOptions, params map[string]string) ([]figma.RenderOptions, error) {
	scales := []float64{}
	if len(opts.Scales) == 0 {
		scales = append(scales, 1.0)
	} else {
		scales = opts.Scales
	}

	// if params["includeID"] == "" {
	// 	params["includeID"] = "false" // default
	// }

	// if params["simplifyStroke"] == "" {
	// 	params["simplifyStroke"] = "false" // default
	// }

	includeID, _ := strconv.ParseBool(params["includeID"])
	simplifyStroke, _ := strconv.ParseBool(params["simplifyStroke"])

	renders := []figma.RenderOptions{}

	for _, scale := range scales {
		renders = append(renders, figma.RenderOptions{
			Scale:             scale,
			Format:            figma.SVGRender,
			SVGIncludeID:      includeID,
			SVGSimplifyStroke: simplifyStroke,
		})
	}

	return nil, nil
}
