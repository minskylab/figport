package figma

import (
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestFigma_downloadFromFigmaRender(t *testing.T) {
	type fields struct {
		config     *viper.Viper
		httpClient *http.Client
	}

	type args struct {
		imageURL string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *os.File
		want1   string
		wantErr bool
	}{
		{
			name:    "Download simple",
			fields:  fields{
				config:     viper.New(),
				httpClient: http.DefaultClient,
			},
			args:    args{
				imageURL: "https://s3-us-west-2.amazonaws.com/figma-alpha-api/img/8cb1/5b2a/7c6a72cb6ea698077132ba0599d839b8",
			},
			want:    nil,
			want1:   "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fig := &Figma{
				config:     tt.fields.config,
				httpClient: tt.fields.httpClient,
			}
			got, got1, err := fig.downloadFromFigmaRender(tt.args.imageURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("downloadFromFigmaRender() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				tt.want, _ = ioutil.TempFile(os.TempDir(), "figport-tests")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("downloadFromFigmaRender() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("downloadFromFigmaRender() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFigma_getFromFigmaFile(t *testing.T) {
	type fields struct {
		config     *viper.Viper
		httpClient *http.Client
	}
	type args struct {
		accessToken string
		fileKey     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fig := &Figma{
				config:     tt.fields.config,
				httpClient: tt.fields.httpClient,
			}
			got, err := fig.getFromFigmaFile(tt.args.accessToken, tt.args.fileKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFromFigmaFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFromFigmaFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFigma_renderImageFromNode(t *testing.T) {
	type fields struct {
		config     *viper.Viper
		httpClient *http.Client
	}
	type args struct {
		accessToken string
		fileKey     string
		nodes       []string
		options     RenderOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Render
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fig := &Figma{
				config:     tt.fields.config,
				httpClient: tt.fields.httpClient,
			}
			got, err := fig.renderImageFromNode(tt.args.accessToken, tt.args.fileKey, tt.args.nodes, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("renderImageFromNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("renderImageFromNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
