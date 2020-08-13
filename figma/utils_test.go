package figma

import (
	"net/http"
	"testing"

	"github.com/spf13/viper"
)

func TestFigma_figmaURI(t *testing.T) {
	type fields struct {
		config     *viper.Viper
		httpClient *http.Client
	}
	type args struct {
		base  string
		paths []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Files API",
			fields:  fields{
				config:     viper.New(),
				httpClient: http.DefaultClient,
			},
			args:    args{
				base: "https://api.figma.com",
				paths: []string{"files", "file_id"},
			},
			want:    "https://api.figma.com/files/file_id",
			wantErr: false,
		},
		{
			name:    "Auth Local API",
			fields:  fields{
				config:     viper.New(),
				httpClient: http.DefaultClient,
			},
			args:    args{
				base: "https://api.figma.com",
				paths: []string{"auth", "local"},
			},
			want:    "https://api.figma.com/auth/local",
			wantErr: false,
		},
		{
			name:    "Invalid Base",
			fields:  fields{
				config:     viper.New(),
				httpClient: http.DefaultClient,
			},
			args:    args{
				base: "://api..figma.",
				paths: []string{"", ""},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fig := &Figma{
				config:     tt.fields.config,
				httpClient: tt.fields.httpClient,
			}

			got, err := fig.figmaURI(tt.args.base, tt.args.paths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("figmaURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("figmaURI() got = %v, want %v", got, tt.want)
			}
		})
	}
}