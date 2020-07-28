package figma

import "testing"

func TestFigma_figmaURI(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name    string
		fig     *Figma
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fig.figmaURI(tt.args.paths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Figma.figmaURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Figma.figmaURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFigma_FigmaURI(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name    string
		fig     *Figma
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fig.FigmaURI(tt.args.paths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Figma.FigmaURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Figma.FigmaURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
