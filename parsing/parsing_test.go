package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestParsing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parsing Test")
}

var _ = Describe("Parsing tag, tagline and make csv", func() {
	Context("Test CSV", func() {
		tags, err := getTaglineAll()
		if err != nil {
			return
		}
		csv := toCSV(tags)

		It("parse success", func() {
			Expect(csv).Should(BeNil())
		})
	})
	Context("Test Parsing", func() {
		taglines, err := getTaglineAll()
		It("get tagline all", func() {
			Expect(taglines[0]).NotTo(BeNil())
			Expect(err).Should(BeNil())
		})
		tag := getTag(taglines[0])
		It("get tag from tagline", func() {
			Expect(tag).NotTo(Equal("error"))
		})
	})
	Context("Test Upload S3", func() {
		It("upload s3 tags csv", func() {
			bucket := "my-note-0.0.1"
			Expect(uploadS3(bucket, "../tags.csv")).Should(BeNil())
		})
	})
})
