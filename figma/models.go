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
	ID                   string          `json:"id"`
	Name                 string          `json:"name"`
	Type                 string          `json:"type"`
	Children             []PurpleChild   `json:"children"`
	BackgroundColor      Color           `json:"backgroundColor"`
	PrototypeStartNodeID *string         `json:"prototypeStartNodeID"`
	PrototypeDevice      PrototypeDevice `json:"prototypeDevice"`
}

// Color details is in Figma official API documentation
type Color struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
	A float64 `json:"a"`
}

// PurpleChild details is in Figma official API documentation
type PurpleChild struct {
	ID                      string              `json:"id"`
	Name                    string              `json:"name"`
	Type                    ChildType           `json:"type"`
	BlendMode               ChildBlendMode      `json:"blendMode"`
	Children                []FluffyChild       `json:"children"`
	AbsoluteBoundingBox     AbsoluteBoundingBox `json:"absoluteBoundingBox"`
	Constraints             Constraints         `json:"constraints"`
	ClipsContent            *bool               `json:"clipsContent,omitempty"`
	Background              []Background        `json:"background"`
	Fills                   []FluffyFill        `json:"fills"`
	Strokes                 []interface{}       `json:"strokes"`
	StrokeWeight            int64               `json:"strokeWeight"`
	StrokeAlign             StrokeAlign         `json:"strokeAlign"`
	BackgroundColor         *Color              `json:"backgroundColor,omitempty"`
	Effects                 []interface{}       `json:"effects"`
	PreserveRatio           *bool               `json:"preserveRatio,omitempty"`
	LayoutGrids             []LayoutGrid        `json:"layoutGrids"`
	ExportSettings          []ExportSetting     `json:"exportSettings"`
	Characters              *string             `json:"characters,omitempty"`
	Style                   *Style              `json:"style,omitempty"`
	CharacterStyleOverrides []interface{}       `json:"characterStyleOverrides"`
}

// AbsoluteBoundingBox details is in Figma official API documentation
type AbsoluteBoundingBox struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// Background details is in Figma official API documentation
type Background struct {
	BlendMode BackgroundBlendMode `json:"blendMode"`
	Type      BackgroundType      `json:"type"`
	Color     Color               `json:"color"`
	Opacity   *float64            `json:"opacity,omitempty"`
}

// FluffyChild details is in Figma official API documentation
type FluffyChild struct {
	ID                      string                              `json:"id"`
	Name                    string                              `json:"name"`
	Type                    ChildType                           `json:"type"`
	Locked                  *bool                               `json:"locked,omitempty"`
	BlendMode               ChildBlendMode                      `json:"blendMode"`
	AbsoluteBoundingBox     AbsoluteBoundingBox                 `json:"absoluteBoundingBox"`
	PreserveRatio           *bool                               `json:"preserveRatio,omitempty"`
	Constraints             Constraints                         `json:"constraints"`
	Fills                   []PurpleFill                        `json:"fills"`
	Strokes                 []Background                        `json:"strokes"`
	StrokeWeight            float64                             `json:"strokeWeight"`
	StrokeAlign             StrokeAlign                         `json:"strokeAlign"`
	Effects                 []interface{}                       `json:"effects"`
	Children                []TentacledChild                    `json:"children"`
	ClipsContent            *bool                               `json:"clipsContent,omitempty"`
	Background              []interface{}                       `json:"background"`
	BackgroundColor         *Color                              `json:"backgroundColor,omitempty"`
	BooleanOperation        *BooleanOperationEnum               `json:"booleanOperation,omitempty"`
	ExportSettings          []ExportSetting                     `json:"exportSettings"`
	Characters              *string                             `json:"characters,omitempty"`
	Style                   *Style                              `json:"style,omitempty"`
	CharacterStyleOverrides []int64                             `json:"characterStyleOverrides"`
	StyleOverrideTable      map[string]FluffyStyleOverrideTable `json:"styleOverrideTable,omitempty"`
	Visible                 *bool                               `json:"visible,omitempty"`
	CornerRadius            *int64                              `json:"cornerRadius,omitempty"`
	RectangleCornerRadii    []int64                             `json:"rectangleCornerRadii"`
	StrokeCap               *Stroke                             `json:"strokeCap,omitempty"`
	StrokeJoin              *Stroke                             `json:"strokeJoin,omitempty"`
	Opacity                 *float64                            `json:"opacity,omitempty"`
}

// TentacledChild details is in Figma official API documentation
type TentacledChild struct {
	ID                      string                              `json:"id"`
	Name                    string                              `json:"name"`
	Type                    ChildType                           `json:"type"`
	BlendMode               ChildBlendMode                      `json:"blendMode"`
	AbsoluteBoundingBox     AbsoluteBoundingBox                 `json:"absoluteBoundingBox"`
	Constraints             Constraints                         `json:"constraints"`
	Fills                   []PurpleFill                        `json:"fills"`
	Strokes                 []Background                        `json:"strokes"`
	StrokeWeight            float64                             `json:"strokeWeight"`
	StrokeAlign             StrokeAlign                         `json:"strokeAlign"`
	Effects                 []Effect                            `json:"effects"`
	PreserveRatio           *bool                               `json:"preserveRatio,omitempty"`
	Children                []StickyChild                       `json:"children"`
	ClipsContent            *bool                               `json:"clipsContent,omitempty"`
	Background              []interface{}                       `json:"background"`
	BackgroundColor         *Color                              `json:"backgroundColor,omitempty"`
	BooleanOperation        *BooleanOperationEnum               `json:"booleanOperation,omitempty"`
	Characters              *string                             `json:"characters,omitempty"`
	Style                   *Style                              `json:"style,omitempty"`
	CharacterStyleOverrides []int64                             `json:"characterStyleOverrides"`
	StyleOverrideTable      map[string]PurpleStyleOverrideTable `json:"styleOverrideTable,omitempty"`
	CornerRadius            *int64                              `json:"cornerRadius,omitempty"`
	StrokeCap               *Stroke                             `json:"strokeCap,omitempty"`
	StrokeJoin              *Stroke                             `json:"strokeJoin,omitempty"`
	RectangleCornerRadii    []int64                             `json:"rectangleCornerRadii"`
	TransitionNodeID        *string                             `json:"transitionNodeID,omitempty"`
	TransitionDuration      *int64                              `json:"transitionDuration,omitempty"`
	TransitionEasing        *string                             `json:"transitionEasing,omitempty"`
	ExportSettings          []ExportSetting                     `json:"exportSettings"`
	Visible                 *bool                               `json:"visible,omitempty"`
}

// StickyChild details is in Figma official API documentation
type StickyChild struct {
	ID                      string                `json:"id"`
	Name                    string                `json:"name"`
	Type                    ChildType             `json:"type"`
	BlendMode               ChildBlendMode        `json:"blendMode"`
	AbsoluteBoundingBox     AbsoluteBoundingBox   `json:"absoluteBoundingBox"`
	Constraints             Constraints           `json:"constraints"`
	Fills                   []PurpleFill          `json:"fills"`
	Strokes                 []Background          `json:"strokes"`
	StrokeWeight            float64               `json:"strokeWeight"`
	StrokeAlign             StrokeAlign           `json:"strokeAlign"`
	Effects                 []interface{}         `json:"effects"`
	PreserveRatio           *bool                 `json:"preserveRatio,omitempty"`
	BooleanOperation        *BooleanOperationEnum `json:"booleanOperation,omitempty"`
	Children                []IndigoChild         `json:"children"`
	ClipsContent            *bool                 `json:"clipsContent,omitempty"`
	Background              []interface{}         `json:"background"`
	BackgroundColor         *Color                `json:"backgroundColor,omitempty"`
	Visible                 *bool                 `json:"visible,omitempty"`
	Characters              *string               `json:"characters,omitempty"`
	Style                   *Style                `json:"style,omitempty"`
	CharacterStyleOverrides []interface{}         `json:"characterStyleOverrides"`
	CornerRadius            *int64                `json:"cornerRadius,omitempty"`
	RectangleCornerRadii    []int64               `json:"rectangleCornerRadii"`
	ExportSettings          []ExportSetting       `json:"exportSettings"`
}

// IndigoChild details is in Figma official API documentation
type IndigoChild struct {
	ID                      string                `json:"id"`
	Name                    string                `json:"name"`
	Type                    ChildType             `json:"type"`
	BlendMode               ChildBlendMode        `json:"blendMode"`
	AbsoluteBoundingBox     AbsoluteBoundingBox   `json:"absoluteBoundingBox"`
	Constraints             Constraints           `json:"constraints"`
	Fills                   []Background          `json:"fills"`
	Strokes                 []Background          `json:"strokes"`
	StrokeWeight            float64               `json:"strokeWeight"`
	StrokeAlign             StrokeAlign           `json:"strokeAlign"`
	Effects                 []interface{}         `json:"effects"`
	PreserveRatio           *bool                 `json:"preserveRatio,omitempty"`
	BooleanOperation        *BooleanOperationEnum `json:"booleanOperation,omitempty"`
	Children                []IndecentChild       `json:"children"`
	ExportSettings          []ExportSetting       `json:"exportSettings"`
	Characters              *string               `json:"characters,omitempty"`
	Style                   *Style                `json:"style,omitempty"`
	CharacterStyleOverrides []interface{}         `json:"characterStyleOverrides"`

	StrokeCap       *Stroke       `json:"strokeCap,omitempty"`
	StrokeJoin      *Stroke       `json:"strokeJoin,omitempty"`
	ClipsContent    *bool         `json:"clipsContent,omitempty"`
	Background      []interface{} `json:"background"`
	BackgroundColor *Color        `json:"backgroundColor,omitempty"`
}

// IndecentChild details is in Figma official API documentation
type IndecentChild struct {
	ID                  string              `json:"id"`
	Name                Name                `json:"name"`
	Type                ChildType           `json:"type"`
	BlendMode           ChildBlendMode      `json:"blendMode"`
	AbsoluteBoundingBox AbsoluteBoundingBox `json:"absoluteBoundingBox"`
	Constraints         Constraints         `json:"constraints"`
	Fills               []Background        `json:"fills"`
	Strokes             []interface{}       `json:"strokes"`
	StrokeWeight        float64             `json:"strokeWeight"`
	StrokeAlign         StrokeAlign         `json:"strokeAlign"`
	Effects             []interface{}       `json:"effects"`
	PreserveRatio       *bool               `json:"preserveRatio,omitempty"`
}

// Constraints details is in Figma official API documentation
type Constraints struct {
	Vertical   Vertical   `json:"vertical"`
	Horizontal Horizontal `json:"horizontal"`
}

// ExportSetting details is in Figma official API documentation
type ExportSetting struct {
	Suffix     string     `json:"suffix"`
	Format     Format     `json:"format"`
	Constraint Constraint `json:"constraint"`
}

// Constraint details is in Figma official API documentation
type Constraint struct {
	Type  ConstraintType `json:"type"`
	Value float64        `json:"value"`
}

// Style details is in Figma official API documentation
type Style struct {
	FontFamily                FontFamily         `json:"fontFamily"`
	FontPostScriptName        FontPostScriptName `json:"fontPostScriptName"`
	FontWeight                int64              `json:"fontWeight"`
	TextAutoResize            *TextAutoResize    `json:"textAutoResize,omitempty"`
	FontSize                  float64            `json:"fontSize"`
	TextAlignHorizontal       Horizontal         `json:"textAlignHorizontal"`
	TextAlignVertical         Vertical           `json:"textAlignVertical"`
	LetterSpacing             float64            `json:"letterSpacing"`
	LineHeightPx              float64            `json:"lineHeightPx"`
	LineHeightPercent         float64            `json:"lineHeightPercent"`
	LineHeightPercentFontSize *float64           `json:"lineHeightPercentFontSize,omitempty"`
	LineHeightUnit            LineHeightUnit     `json:"lineHeightUnit"`
	Hyperlink                 *StyleHyperlink    `json:"hyperlink,omitempty"`
	TextDecoration            *string            `json:"textDecoration,omitempty"`
}

// StyleHyperlink details is in Figma official API documentation
type StyleHyperlink struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// PurpleFill details is in Figma official API documentation
type PurpleFill struct {
	BlendMode               BackgroundBlendMode      `json:"blendMode"`
	Type                    BackgroundType           `json:"type"`
	Color                   *Color                   `json:"color,omitempty"`
	ScaleMode               *ScaleMode               `json:"scaleMode,omitempty"`
	ImageRef                *string                  `json:"imageRef,omitempty"`
	Opacity                 *float64                 `json:"opacity,omitempty"`
	ImageTransform          [][]float64              `json:"imageTransform"`
	GradientHandlePositions []GradientHandlePosition `json:"gradientHandlePositions"`
	GradientStops           []GradientStop           `json:"gradientStops"`
	Visible                 *bool                    `json:"visible,omitempty"`
}

// GradientHandlePosition details is in Figma official API documentation
type GradientHandlePosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// GradientStop details is in Figma official API documentation
type GradientStop struct {
	Color    Color `json:"color"`
	Position int64 `json:"position"`
}

// Effect details is in Figma official API documentation
type Effect struct {
	Type      string                 `json:"type"`
	Visible   bool                   `json:"visible"`
	Color     Color                  `json:"color"`
	BlendMode BackgroundBlendMode    `json:"blendMode"`
	Offset    GradientHandlePosition `json:"offset"`
	Radius    int64                  `json:"radius"`
}

// PurpleStyleOverrideTable details is in Figma official API documentation
type PurpleStyleOverrideTable struct {
	FontSize float64      `json:"fontSize"`
	Fills    []Background `json:"fills"`
}

// FluffyStyleOverrideTable details is in Figma official API documentation
type FluffyStyleOverrideTable struct {
	Fills          []Background                 `json:"fills"`
	Hyperlink      *StyleOverrideTableHyperlink `json:"hyperlink,omitempty"`
	TextDecoration *string                      `json:"textDecoration,omitempty"`
}

// StyleOverrideTableHyperlink details is in Figma official API documentation
type StyleOverrideTableHyperlink struct {
	Type *string `json:"type,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// FluffyFill details is in Figma official API documentation
type FluffyFill struct {
	BlendMode BackgroundBlendMode `json:"blendMode"`
	Type      BackgroundType      `json:"type"`
	Color     *Color              `json:"color,omitempty"`
	ScaleMode *ScaleMode          `json:"scaleMode,omitempty"`
	ImageRef  *string             `json:"imageRef,omitempty"`
}

// LayoutGrid details is in Figma official API documentation
type LayoutGrid struct {
	Pattern     Pattern   `json:"pattern"`
	SectionSize float64   `json:"sectionSize"`
	Visible     bool      `json:"visible"`
	Color       Color     `json:"color"`
	Alignment   ScaleMode `json:"alignment"`
	GutterSize  int64     `json:"gutterSize"`
	Offset      int64     `json:"offset"`
	Count       int64     `json:"count"`
}

// PrototypeDevice details is in Figma official API documentation
type PrototypeDevice struct {
	// Type details in Figma API Documentation
	Type string `json:"type"`
	// Rotation details in Figma API Documentation
	Rotation string `json:"rotation"`
}

// BackgroundBlendMode details is in Figma official API documentation
type BackgroundBlendMode string

const (
	// Normal details in Figma API Documentation
	Normal BackgroundBlendMode = "NORMAL"
)

// BackgroundType details is in Figma official API documentation
type BackgroundType string

const (
	// GradientLinear details in Figma API Documentation
	GradientLinear BackgroundType = "GRADIENT_LINEAR"
	// Image details in Figma API Documentation
	Image BackgroundType = "IMAGE"
	// Solid details in Figma API Documentation
	Solid BackgroundType = "SOLID"
)

// ChildBlendMode details is in Figma official API documentation
type ChildBlendMode string

const (
	// Darken details in Figma API Documentation
	Darken ChildBlendMode = "DARKEN"
	// PassThrough details in Figma API Documentation
	PassThrough ChildBlendMode = "PASS_THROUGH"
	// Saturation details in Figma API Documentation
	Saturation ChildBlendMode = "SATURATION"
)

// BooleanOperationEnum details is in Figma official API documentation
type BooleanOperationEnum string

const (
	// Intersect details in Figma API Documentation
	Intersect BooleanOperationEnum = "INTERSECT"
	// Subtract details in Figma API Documentation
	Subtract BooleanOperationEnum = "SUBTRACT"
	// Union details in Figma API Documentation
	Union BooleanOperationEnum = "UNION"
)

// Horizontal details is in Figma official API documentation
type Horizontal string

const (
	// HorizontalCENTER details in Figma API Documentation
	HorizontalCENTER Horizontal = "CENTER"
	// Left details in Figma API Documentation
	Left Horizontal = "LEFT"
	// Right details in Figma API Documentation
	Right Horizontal = "RIGHT"
)

// Vertical details is in Figma official API documentation
type Vertical string

const (
	// Top details in Figma API Documentation
	Top Vertical = "TOP"
	// VerticalCENTER details in Figma API Documentation
	VerticalCENTER Vertical = "CENTER"
)

// Name details is in Figma official API documentation
type Name string

const (
	// Ellipse27 details in Figma API Documentation
	Ellipse27 Name = "Ellipse 27"
	// Ellipse34 details in Figma API Documentation
	Ellipse34 Name = "Ellipse 34"
	// NameUnion details in Figma API Documentation
	NameUnion Name = "Union"
)

// StrokeAlign details is in Figma official API documentation
type StrokeAlign string

const (
	// Inside details in Figma API Documentation
	Inside StrokeAlign = "INSIDE"
	// Outside details in Figma API Documentation
	Outside StrokeAlign = "OUTSIDE"
	// StrokeAlignCENTER details in Figma API Documentation
	StrokeAlignCENTER StrokeAlign = "CENTER"
)

// ChildType details is in Figma official API documentation
type ChildType string

const (
	// BooleanOperation details in Figma API Documentation
	BooleanOperation ChildType = "BOOLEAN_OPERATION"
	// Ellipse details in Figma API Documentation
	Ellipse ChildType = "ELLIPSE"
	// Frame details in Figma API Documentation
	Frame ChildType = "FRAME"
	// Group details in Figma API Documentation
	Group ChildType = "GROUP"
	// Line details in Figma API Documentation
	Line ChildType = "LINE"
	// Rectangle details in Figma API Documentation
	Rectangle ChildType = "RECTANGLE"
	// Text details in Figma API Documentation
	Text ChildType = "TEXT"
	// Vector details in Figma API Documentation
	Vector ChildType = "VECTOR"
)

// ConstraintType details is in Figma official API documentation
type ConstraintType string

const (
	// Scale details in Figma API Documentation
	Scale ConstraintType = "SCALE"
)

// Format details is in Figma official API documentation
type Format string

const (
	// PNG details in Figma API Documentation
	PNG Format = "PNG"
	// SVG details in Figma API Documentation
	SVG Format = "SVG"
)

// Stroke details is in Figma official API documentation
type Stroke string

const (
	// Round details in Figma API Documentation
	Round Stroke = "ROUND"
	// TriangleArrow details in Figma API Documentation
	TriangleArrow Stroke = "TRIANGLE_ARROW"
)

// FontFamily details is in Figma official API documentation
type FontFamily string

const (
	// Barlow details in Figma API Documentation
	Barlow FontFamily = "Barlow"
	// Heebo details in Figma API Documentation
	Heebo FontFamily = "Heebo"
	// OpenSans details in Figma API Documentation
	OpenSans FontFamily = "Open Sans"
	// Rubik details in Figma API Documentation
	Rubik FontFamily = "Rubik"
)

// FontPostScriptName details is in Figma official API documentation
type FontPostScriptName string

const (
	// BarlowMedium details in Figma API Documentation
	BarlowMedium FontPostScriptName = "Barlow-Medium"
	// HeeboMedium details in Figma API Documentation
	HeeboMedium FontPostScriptName = "Heebo-Medium"
	// OpenSansBold details in Figma API Documentation
	OpenSansBold FontPostScriptName = "OpenSans-Bold"
	// OpenSansRegular details in Figma API Documentation
	OpenSansRegular FontPostScriptName = "OpenSans-Regular"
	// OpenSansSemiBold details in Figma API Documentation
	OpenSansSemiBold FontPostScriptName = "OpenSans-SemiBold"
	// RubikBold details in Figma API Documentation
	RubikBold FontPostScriptName = "Rubik-Bold"
	// RubikMedium details in Figma API Documentation
	RubikMedium FontPostScriptName = "Rubik-Medium"
	// RubikRegular details in Figma API Documentation
	RubikRegular FontPostScriptName = "Rubik-Regular"
)

// LineHeightUnit details is in Figma official API documentation
type LineHeightUnit string

const (
	// FontSize details in Figma API Documentation
	FontSize LineHeightUnit = "FONT_SIZE_%"
	// Intrinsic details in Figma API Documentation
	Intrinsic LineHeightUnit = "INTRINSIC_%"
)

// TextAutoResize details is in Figma official API documentation
type TextAutoResize string

const (
	// Height details in Figma API Documentation
	Height TextAutoResize = "HEIGHT"
	// WidthAndHeight details in Figma API Documentation
	WidthAndHeight TextAutoResize = "WIDTH_AND_HEIGHT"
)

// ScaleMode details is in Figma official API documentation
type ScaleMode string

const (
	// Fill details in Figma API Documentation
	Fill ScaleMode = "FILL"
	// Stretch details in Figma API Documentation
	Stretch ScaleMode = "STRETCH"
)

// Pattern details is in Figma official API documentation
type Pattern string

const (
	// Columns details in Figma API Documentation
	Columns Pattern = "COLUMNS"
)
