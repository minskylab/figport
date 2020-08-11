package figport

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (fig *Figport) connectS3() error {
	s3Options := getMinioClientOptionsFromConfig(fig.config)

	logrus.WithFields(logrus.Fields{
		"s3Endpoint": s3Options.Endpoint,
	}).Info("establishment s3 connection")

	client, err := minio.New(s3Options.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s3Options.AccessKeyID, s3Options.SecretKey, ""),
		Region: s3Options.Region,
		Secure: s3Options.UseSSL,
	})

	if err != nil {
		return errors.WithStack(err)
	}

	fig.s3Client = client
	fig.s3Options = s3Options

	return nil
}
