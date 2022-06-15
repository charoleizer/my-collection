package domain_test

import (
	"fmt"
	"runtime"
	"time"

	"github.com/charoleizer/my-collection/src/domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
)

var _ = Describe("Domain", func() {

	Context("When HelloSekai return a string", func() {

		It("Should return 'Hello, 世界'", func() {
			Expect(domain.HelloSekai()).To(Equal("Hello, 世界"))
		})

		It("Should perform in less than 500ms", func() {
			experiment := gmeasure.NewExperiment("HelloSekai Performance")
			experiment.SampleDuration("HelloSekai", func(_ int) {
				domain.HelloSekai()
			}, gmeasure.SamplingConfig{N: 1})
			Expect(experiment.GetStats("HelloSekai").DurationFor(gmeasure.StatMedian)).Should(BeNumerically("<", 500*time.Millisecond))

		})

		It("Should not exceeding 1 mb of memory", func() {
			experiment := gmeasure.NewExperiment("exploring the HelloSekai")

			experiment.RecordNote("HelloSekai Properties")
			experiment.Sample(func(idx int) {
				stopwatch := experiment.NewStopwatch()
				domain.HelloSekai()
				stopwatch.Record("HelloSekai Runtime", gmeasure.Style(""), gmeasure.Precision(time.Millisecond))

				var m runtime.MemStats
				runtime.ReadMemStats(&m)

				experiment.RecordValue("HelloSekai Memory Usage", float64(m.Alloc)/1024/1024, gmeasure.Style(""), gmeasure.Precision(3), gmeasure.Units("MB"), gmeasure.Annotation(fmt.Sprintf("%d", idx)))
			}, gmeasure.SamplingConfig{N: 1000, NumParallel: 4})

			memoryStats := experiment.GetStats("HelloSekai Memory Usage")
			minMemory := memoryStats.ValueFor(gmeasure.StatMin)
			maxMemory := memoryStats.ValueFor(gmeasure.StatMax)

			Expect(maxMemory-minMemory).To(BeNumerically("<=", 1), "Should not see memory fluctuations exceeding 1 mb")
		})
	})

})
