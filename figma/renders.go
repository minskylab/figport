package figma

// Render is the expected response from the figma /v1/images/:key endpoint
type Render struct {
	Err    string            `json:"err"`
	Images map[string]string `json:"images"`
	Status float64           `json:"status"`
}

// RenderFormat determine the export type of your Figma Node
type RenderFormat string

const jpg RenderFormat = "jpg"
const png RenderFormat = "png"
const svg RenderFormat = "svg"
const pdf RenderFormat = "pdf"

// RenderOptions wraps the options to specify your export (render) task
type RenderOptions struct {
	Scale             float64
	Format            RenderFormat
	SVGIncludeID      bool
	SVGSimplifyStroke bool
	UseAbsoluteBounds bool
	Version           string
}
