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

	Context("When HelloFigures return an struct", func() {

		It("Should return 'Luffy' with ID 1", func() {
			figure := domain.HelloFigures()
			Expect(figure.ID).To(Equal("1"))
			Expect(figure.Name).To(Equal("Luffy"))
		})

		It("Should perform in less than 500ms", func() {
			experiment := gmeasure.NewExperiment("HelloFigures Performance")
			experiment.SampleDuration("HelloFigures", func(_ int) {
				domain.HelloFigures()
			}, gmeasure.SamplingConfig{N: 1})
			Expect(experiment.GetStats("HelloFigures").DurationFor(gmeasure.StatMedian)).Should(BeNumerically("<", 500*time.Millisecond))

		})

		It("Should not exceeding 1 mb of memory", func() {
			experiment := gmeasure.NewExperiment("exploring the HelloFigures")

			experiment.RecordNote("HelloFigures Properties")
			experiment.Sample(func(idx int) {
				stopwatch := experiment.NewStopwatch()
				domain.HelloFigures()
				stopwatch.Record("HelloFigures Runtime", gmeasure.Style(""), gmeasure.Precision(time.Millisecond))

				var m runtime.MemStats
				runtime.ReadMemStats(&m)

				experiment.RecordValue("HelloFigures Memory Usage", float64(m.Alloc)/1024/1024, gmeasure.Style(""), gmeasure.Precision(3), gmeasure.Units("MB"), gmeasure.Annotation(fmt.Sprintf("%d", idx)))
			}, gmeasure.SamplingConfig{N: 1000, NumParallel: 4})

			memoryStats := experiment.GetStats("HelloFigures Memory Usage")
			minMemory := memoryStats.ValueFor(gmeasure.StatMin)
			maxMemory := memoryStats.ValueFor(gmeasure.StatMax)

			Expect(maxMemory-minMemory).To(BeNumerically("<=", 1), "Should not see memory fluctuations exceeding 1 mb")
		})
	})

})
