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
	Context("Test get telegram", func() {
		It("pong", func() {
			//Expect(getTelegram()).Should(BeNil())
		})
	})
	Context("Test call SQS", func() {
		//message := textToString("../recommend.txt")
		message, err := ioutil.ReadFile("../recommend.txt")
		It("get queue", func() {
			Expect(sendSqs(string(message))).Should(BeNil())
			Expect(err).Should(BeNil())
		})
	})
})
