package exporting

// ExportNodeOptions is the result of parse a node name from Figma
type ExportNodeOptions struct {
	Raw      string
	Path     string
	Filename string
	Format   string
	Mods     string
	Scale    string
}
