package flusher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tchorzewski1991/flusher/flusher"
)

var _ = Describe("Agent", func() {
	var agent *flusher.Agent

	BeforeEach(func() {
		agent = flusher.NewAgent()
	})

	Context("Start", func() {
		It("starts agent", func() {
			Expect(agent.Start()).NotTo(HaveOccurred())
			Expect(agent.Active()).To(BeTrue())
		})
	})

	Context("Stop", func() {
		It("stops agent", func() {
			Expect(agent.Stop()).NotTo(HaveOccurred())
			Expect(agent.Active()).To(BeFalse())
		})
	})
})
