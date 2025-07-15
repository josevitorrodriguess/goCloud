package aws

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	Bucket string
	Region string
}

func NewS3Client(ctx context.Context, cfg S3Config) (*s3.Client, error) {
	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(cfg.Region),
	)
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(awsCfg), nil
}

func GetS3ConfigFromEnv() S3Config {
	return S3Config{
		Bucket: os.Getenv("AWS_S3_BUCKET"),
		Region: os.Getenv("AWS_REGION"),
	}
}
