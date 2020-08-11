package figma

// File details is in Figma official API documentation
type File struct {
	Document      Document             `json:"document"`
	Components    map[string]Component `json:"components"`
	SchemaVersion int64                `json:"schemaVersion"`
	Styles        Styles               `json:"styles"`
	Name          string               `json:"name"`
	LastModified  string               `json:"lastModified"`
	ThumbnailURL  string               `json:"thumbnailUrl"`
	Version       string               `json:"version"`
	Role          string               `json:"role"`
}

// Component details is in Figma official API documentation
type Component struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Styles details is in Figma official API documentation
type Styles struct{}

// Document details is in Figma official API documentation
type Document struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Children []Canvas `json:"children"`
}

// Canvas details is in Figma official API documentation
type Canvas struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Children []Canvas `json:"children"`
}
