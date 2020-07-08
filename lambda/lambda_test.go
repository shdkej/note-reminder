package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestParsing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cloud Test")
}

var _ = Describe("AWS execute test", func() {
	Context("Test Upload S3", func() {
		It("upload s3 tags csv", func() {
			bucket := "my-note-0.0.1"
			Expect(uploadS3(bucket, "../result/tags.csv")).Should(BeNil())
		})
	})
})
