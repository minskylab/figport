package figport

import (
	"context"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minskylab/figport/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type s3storageOptions struct {
	Endpoint    string
	AccessKeyID string
	SecretKey   string
	UseSSL      bool
	Region      string
}

func getMinioClientOptionsFromConfig(conf *viper.Viper) *s3storageOptions {
	s3Endpoint := conf.GetString(config.S3Endpoint)
	s3AccessKeyID := conf.GetString(config.S3AccessKeyID)
	s3SecretKey := conf.GetString(config.S3SecretKey)
	s3UseSSL := conf.GetBool(config.S3UseSSL)
	s3Region := conf.GetString(config.S3Region)

	return &s3storageOptions{
		Endpoint:    s3Endpoint,
		AccessKeyID: s3AccessKeyID,
		SecretKey:   s3SecretKey,
		UseSSL:      s3UseSSL,
		Region:      s3Region,
	}
}

func (fig *Figport) saveAsset(ctx context.Context, path string, contentType string, file *os.File) (interface{}, error) {
	bucket := fig.config.GetString(config.S3Bucket)

	bucketIsOk, err := fig.s3Client.BucketExists(ctx, bucket)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !bucketIsOk {
		return nil, errors.New("your bucket not exist, aborting saving operation")
	}

	info, err := fig.s3Client.FPutObject(ctx, bucket, path, file.Name(), minio.PutObjectOptions{
		UserMetadata: map[string]string{
			"uploader": "figport agent",
		},
		ContentType: contentType,
		// TODO: Add more metadata for plugins
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	logrus.WithField("location", info.Bucket).Info("saved asset to s3 storage")

	return nil, nil
}
