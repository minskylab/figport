package figport

import (
	"context"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minskylab/figport/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func getMinioClientFromConfig(conf *viper.Viper) *minio.Client {
	s3Endpoint := conf.GetString(config.S3Endpoint)
	s3AccessKeyID := conf.GetString(config.S3AccessKeyID)
	s3SecretKey := conf.GetString(config.S3SecretKey)
	s3UseSSL := conf.GetBool(config.S3UseSSL)
	s3Region := conf.GetString(config.S3Region)

	client, err := minio.New(s3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s3AccessKeyID, s3SecretKey, ""),
		Region: s3Region,
		Secure: s3UseSSL,
	})

	if err != nil {
		logrus.Panic(errors.WithStack(err))
	}

	return client
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

	info, err := fig.s3Client.PutObject(ctx, bucket, path, file, -1, minio.PutObjectOptions{
		UserMetadata: map[string]string{
			"uploader": "figport agent",
		},
		ContentType: contentType,
		// TODO: Add more metadata for plugins
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	logrus.WithField("location", info.Location).Info("saved asset to s3 storage")

	return nil, nil
}
