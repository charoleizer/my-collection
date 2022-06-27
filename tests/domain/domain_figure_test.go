package domain_test

import (
	"fmt"
	"runtime"
	"time"

	"github.com/charoleizer/my-collection/src/domain"
	"github.com/charoleizer/my-collection/src/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
)

var _ = Describe("Figure Domain", func() {

	Context("For 'IsLuffy' validation", func() {

		It("Should return True if Figure's name is Luffy", func() {
			figure := models.Figures{"1", "Luffy"}
			Expect(domain.IsLuffy(figure)).To(BeTrue())
		})

		It("Should perform in less than 500ms", func() {
			experiment := gmeasure.NewExperiment("IsLuffy Performance")
			experiment.SampleDuration("IsLuffy", func(_ int) {
				domain.IsLuffy(models.Figures{})
			}, gmeasure.SamplingConfig{N: 1})
			Expect(experiment.GetStats("IsLuffy").DurationFor(gmeasure.StatMedian)).Should(BeNumerically("<", 500*time.Millisecond))

		})

		It("Should not exceeding 1 mb of memory", func() {
			experiment := gmeasure.NewExperiment("exploring the IsLuffy")

			experiment.RecordNote("HelloFigures Properties")
			experiment.Sample(func(idx int) {
				stopwatch := experiment.NewStopwatch()
				domain.IsLuffy(models.Figures{})
				stopwatch.Record("IsLuffy Runtime", gmeasure.Style(""), gmeasure.Precision(time.Millisecond))

				var m runtime.MemStats
				runtime.ReadMemStats(&m)

				experiment.RecordValue("IsLuffy Memory Usage", float64(m.Alloc)/1024/1024, gmeasure.Style(""), gmeasure.Precision(3), gmeasure.Units("MB"), gmeasure.Annotation(fmt.Sprintf("%d", idx)))
			}, gmeasure.SamplingConfig{N: 1000, NumParallel: 4})

			memoryStats := experiment.GetStats("IsLuffy Memory Usage")
			minMemory := memoryStats.ValueFor(gmeasure.StatMin)
			maxMemory := memoryStats.ValueFor(gmeasure.StatMax)

			Expect(maxMemory-minMemory).To(BeNumerically("<=", 1), "Should not see memory fluctuations exceeding 1 mb")
		})
	})

})
