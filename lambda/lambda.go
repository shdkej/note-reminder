package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

func uploadS3(bucket string, filepath string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	svc := s3manager.NewUploader(sess)

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	log.Println("File Upload: ", filepath)

	_, err = svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("note/" + filepath),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("File Upload Complete")
	return nil
}
