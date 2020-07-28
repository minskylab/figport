package figma

// Render is the expected response from the figma /v1/images/:key endpoint
type Render struct {
	Err    string            `json:"err"`
	Images map[string]string `json:"images"`
	Status float64           `json:"status"`
}

// RenderFormat determine the export type of your Figma Node
type RenderFormat string

// JPGRender is a type of Figma render
const JPGRender RenderFormat = "jpg"

// PNGRender is a type of Figma render
const PNGRender RenderFormat = "png"

// SVGRender is a type of Figma render
const SVGRender RenderFormat = "svg"

// PDFRender is a type of Figma render
const PDFRender RenderFormat = "pdf"

// RenderOptions wraps the options to specify your export (render) task
type RenderOptions struct {
	Scale             float64
	Format            RenderFormat
	SVGIncludeID      bool
	SVGSimplifyStroke bool
	UseAbsoluteBounds bool
	Version           string
}
