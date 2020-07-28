package figma

import (
	"os"
	"reflect"
	"testing"
)

func TestFigma_ObtainImage(t *testing.T) {
	type args struct {
		accessToken string
		fileKey     string
		nodeID      string
		options     RenderOptions
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
			got, got1, err := tt.fig.ObtainImage(tt.args.accessToken, tt.args.fileKey, tt.args.nodeID, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Figma.ObtainImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Figma.ObtainImage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Figma.ObtainImage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
