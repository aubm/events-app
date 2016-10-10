package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Main", []Reporter{junitReporter})
}

var _ = Describe("Events", func() {

	var events []Event

	BeforeEach(func() {
		events = getEvents()
	})

	It("should have 2 events", func() {
		Expect(len(events)).To(Equal(2))
	})

	It("should have a first event named ToulouseJS", func() {
		Expect(events[0].Name).To(Equal("ToulouseJS"))
	})

})
