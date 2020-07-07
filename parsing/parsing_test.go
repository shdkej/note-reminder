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
		if wikiDir == "" {
			wikiDir = "/home/sh/vimwiki"
		}
		filename := wikiDir + "/Data_Structure.md"
		taglines, err := getTaglines(filename)
		It("get taglines", func() {
			Expect(taglines).NotTo(BeNil())
			Expect(err).Should(BeNil())
		})
		tag := getTag(taglines[0])
		It("get tag from tagline", func() {
			Expect(tag).NotTo(Equal("error"))
		})
	})
})
