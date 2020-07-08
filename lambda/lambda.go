package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
	"path/filepath"
)

func uploadS3(bucket string, path string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	svc := s3manager.NewUploader(sess)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	log.Println("File Upload: ", path)

	filename := filepath.Base(path)
	_, err = svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("File Upload Complete")
	return nil
}
