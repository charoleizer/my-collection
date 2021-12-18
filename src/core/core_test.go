package core_test

import (
	"time"

	"github.com/charoleizer/my-collection/src/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
)

var _ = Describe("Core", func() {

	It("HelloSekai should return 'Hello, 世界'", func() {
		Expect(core.HelloSekai()).To(Equal("Hello, 世界"))
	})

	It("HelloSekai should perform in less than 500ms", func() {
		experiment := gmeasure.NewExperiment("HelloSekai Performance")
		experiment.SampleDuration("HelloSekai", func(_ int) {
			core.HelloSekai()
		}, gmeasure.SamplingConfig{N: 1})
		Expect(experiment.GetStats("HelloSekai").DurationFor(gmeasure.StatMedian)).Should(BeNumerically("<", 500*time.Millisecond))

	})

})
