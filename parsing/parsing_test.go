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
		tags, err := getTagAll()
		It("get tag all", func() {
			Expect(tags[0]).Should(Equal(tag))
			Expect(err).Should(BeNil())
		})
	})
	Context("Test Render index.html", func() {
		recommend := "../recommend.txt"
		recommend_string, err := textToString(recommend)
		It("get tag all", func() {
			Expect(toHTML(recommend_string)).Should(BeNil())
			Expect(err).Should(BeNil())
		})
	})
})
