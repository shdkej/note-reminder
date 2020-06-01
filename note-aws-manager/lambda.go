package main

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"net/http"
	"os"
)

func getTelegram() error {
	url := "https://j1pzc4zmw9.execute-api.eu-central-1.amazonaws.com/dev/send-telegram"
	var data = []byte(`{"message": {"text":"testda!"}}`)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

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

func sendSqs(message string) error {
	if message == "" {
		message = "message is empty"
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")},
	)
	svc := sqs.New(sess)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String("MyQueue"),
	})
	if err != nil {
		return err
	}
	qURL := *result.QueueUrl
	log.Println("Success", qURL)

	sendMessage, err := svc.SendMessage(&sqs.SendMessageInput{
		//DelaySeconds: aws.Int64(10),
		//MessageAttributes: map[string]*sqs.MessageAttributeValue{
		//	"Message": &sqs.MessageAttributeValue{
		//		DataType:    aws.String("String"),
		//		StringValue: aws.String("gogogogogogogo"),
		//	},
		//},
		MessageBody: aws.String(message),
		QueueUrl:    &qURL,
	})
	if err != nil {
		return err
	}

	log.Println("Send Success", *sendMessage.MessageId)

	return nil
}
