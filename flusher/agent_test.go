package flusher_test

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tchorzewski1991/flusher/flusher"
)

var _ = Describe("Agent", func() {
	var agent *flusher.Agent
	var err error

	BeforeEach(func() {
		agent, err = flusher.NewAgent()
	})

	Context("NewAgent", func() {
		It("should propagate opt err", func() {
			agent, err = flusher.NewAgent(func(_ *flusher.Agent) error {
				return errors.New("opt err")
			})
			Expect(agent).To(BeNil())
			Expect(err.Error()).To(Equal("opt err"))
		})
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
