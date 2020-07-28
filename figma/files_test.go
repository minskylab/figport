package figma

import (
	"os"
	"reflect"
	"testing"
)

func TestFigma_getFromFigmaFile(t *testing.T) {
	type args struct {
		accessToken string
		fileKey     string
		nodes       []string
	}
	tests := []struct {
		name    string
		fig     *Figma
		args    args
		want    *File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fig.getFromFigmaFile(tt.args.accessToken, tt.args.fileKey, tt.args.nodes...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Figma.getFromFigmaFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Figma.getFromFigmaFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFigma_renderImageFromNode(t *testing.T) {
	type args struct {
		accessToken string
		fileKey     string
		nodes       []string
		options     RenderOptions
	}
	tests := []struct {
		name    string
		fig     *Figma
		args    args
		want    *Render
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fig.renderImageFromNode(tt.args.accessToken, tt.args.fileKey, tt.args.nodes, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Figma.renderImageFromNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Figma.renderImageFromNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFigma_downloadFromFigmaRender(t *testing.T) {
	type args struct {
		imageURL string
	}
	tests := []struct {
		name    string
		fig     *Figma
		args    args
		want    *os.File
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.fig.downloadFromFigmaRender(tt.args.imageURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Figma.downloadFromFigmaRender() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Figma.downloadFromFigmaRender() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Figma.downloadFromFigmaRender() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
