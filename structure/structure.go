package structure

import (
	"regexp"
	"strings"

	"github.com/minskylab/figport/figma"
)

// ExtractStructureFromFigmaFile ...
func ExtractStructureFromFigmaFile(figmaFile *figma.File) []string {
	form := regexp.MustCompile("[a-z_]+/[a-z_/]*")

	structure := []string{}

	for _, f := range figmaFile.Document.Children {
		for _, frame := range f.Children {
			if frame.Type != "FRAME" {
				continue
			}

			if strings.HasPrefix(frame.Name, "_") {
				continue
			}

			if !form.MatchString(frame.Name) {
				continue
			}

			structure = append(structure, strings.TrimSpace(frame.Name))
		}
	}

	return structure
}
