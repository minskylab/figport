package figport

import (
	"github.com/minskylab/figport/exporting"
	"github.com/minskylab/figport/figma"
)

// Mod is a necessary interface to implement a figma render mod,
// TODO: Improve that
type Mod interface {
	Name() string
	Params() map[string]string
	Process(opts exporting.ExportNodeOptions, params map[string]string) ([]figma.RenderOptions, error)
}
