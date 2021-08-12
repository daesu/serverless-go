package infrastructure

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetAWSSession returns an AWS session initialized
// with the REGION environment variable.
func GetAWSSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")),
	})

	if err != nil {
		panic("Need valid AWS session to operate")
	}

	return sess
}

// GetS3Client takes an and AWS session and returns an s3 client
func GetS3Client(sess client.ConfigProvider) *s3.S3 {
	svc := s3.New(sess)
	return svc
}
