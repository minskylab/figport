package mods

import (
	"github.com/minskylab/figport/exporting"
	"github.com/minskylab/figport/figma"
)

type imageFormat struct {
	Name string
}

type SVGMod struct{ imageFormat }

func (mod *SVGMod) Name() string {
	return "svg"
}

func (mod *SVGMod) Params() map[string]string {
	return map[string]string{}
}

func (mod *SVGMod) Process(opts *exporting.ExportNodeOptions, params map[string]string) ([]figma.RenderOptions, error) {
	return nil, nil
}
