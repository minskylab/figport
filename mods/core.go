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

	return renders, nil
}

type PNGMod struct{}

func (mod *PNGMod) Name() string {
	return "png"
}

func (mod *PNGMod) Params() map[string]string {
	return map[string]string{}
}

func (mod *PNGMod) Process(opts exporting.ExportNodeOptions, params map[string]string) ([]figma.RenderOptions, error) {
	scales := []float64{}
	if len(opts.Scales) == 0 {
		scales = append(scales, 1.0)
	} else {
		scales = opts.Scales
	}

	renders := []figma.RenderOptions{}

	for _, scale := range scales {
		renders = append(renders, figma.RenderOptions{
			Scale:  scale,
			Format: figma.PNGRender,
		})
	}

	return renders, nil
}

type JPGMod struct{}

func (mod *JPGMod) Name() string {
	return "jpg"
}

func (mod *JPGMod) Params() map[string]string {
	return map[string]string{}
}

func (mod *JPGMod) Process(opts exporting.ExportNodeOptions, params map[string]string) ([]figma.RenderOptions, error) {
	scales := []float64{}
	if len(opts.Scales) == 0 {
		scales = append(scales, 1.0)
	} else {
		scales = opts.Scales
	}

	renders := []figma.RenderOptions{}

	for _, scale := range scales {
		renders = append(renders, figma.RenderOptions{
			Scale:  scale,
			Format: figma.JPGRender,
		})
	}

	return renders, nil
}

type PDFMod struct{}

func (mod *PDFMod) Name() string {
	return "jpg"
}

func (mod *PDFMod) Params() map[string]string {
	return map[string]string{}
}

func (mod *PDFMod) Process(opts exporting.ExportNodeOptions, params map[string]string) ([]figma.RenderOptions, error) {
	scales := []float64{}
	if len(opts.Scales) == 0 {
		scales = append(scales, 1.0)
	} else {
		scales = opts.Scales
	}

	renders := []figma.RenderOptions{}

	for _, scale := range scales {
		renders = append(renders, figma.RenderOptions{
			Scale:  scale,
			Format: figma.PDFRender,
		})
	}

	return renders, nil
}
