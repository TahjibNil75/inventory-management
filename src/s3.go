package src

import (
	"bytes"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func uploadToS3(data []byte, key string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})
	if err != nil {
		return err
	}

	s3Client := s3.New(sess)

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(os.Getenv("BUCKET_NAME")),
		Key:           aws.String(key),
		Body:          bytes.NewReader(data),
		ContentLength: aws.Int64(int64(len(data))),
		ContentType:   aws.String("image/png"),
		ACL:           aws.String("public-read"),
	})
	return err
}

func CheckIfAssetExistsinS3(assetTag string) bool {
	// Correct the S3 key for the asset's QR code.
	key := "qrcode/" + assetTag + ".png" // Fix the typo here
	// Aws session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})
	if err != nil {
		return false
	}
	s3Client := s3.New(sess)

	_, err = s3Client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(key),
	})
	return err == nil
}

func DeleteFromS3(key string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})
	if err != nil {
		return err
	}

	s3Client := s3.New(sess)
	_, err = s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(key),
	})
	return err
}

