package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"testing"
)

func TestParsing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parsing Test")
}

var _ = Describe("AWS execute test", func() {
	path := os.Getenv("CSV_PATH")
	if path == "" {
		path = "../result/tags.csv"
	}
	Context("Test Upload S3", func() {
		It("upload s3 tags csv", func() {
			bucket := "my-note-0.0.1"
			Expect(uploadS3(bucket, "")).Should(BeNil())
		})
	})
})
