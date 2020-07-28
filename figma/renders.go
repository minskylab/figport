package figma

// Render is the expected response from the figma /v1/images/:key endpoint
type Render struct {
	Err    string            `json:"err"`
	Images map[string]string `json:"images"`
	Status float64           `json:"status"`
}
