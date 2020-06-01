package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"testing"
)

func TestAWS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parsing AWS")
}

var _ = Describe("Get AWS", func() {
	filepath := "../recommend.txt"
	Context("Test get telegram", func() {
		It("pong", func() {
			//Expect(getTelegram()).Should(BeNil())
		})
	})
	Context("Test call SQS", func() {
		//message := textToString("../recommend.txt")
		message, err := ioutil.ReadFile(filepath)
		It("get queue", func() {
			Expect(sendSqs(string(message))).Should(BeNil())
			Expect(err).Should(BeNil())
		})
	})
	Context("Test Upload S3", func() {
		bucket := "s3-web-hosting-test-0.0.1"
		It("upload s3 recommend.txt", func() {
			Expect(uploadS3(bucket, filepath)).Should(BeNil())
		})
		bucket = "my-note-0.0.1"
		It("upload s3 tags csv", func() {
			Expect(uploadS3(bucket, "../tags.csv")).Should(BeNil())
		})
	})
})
