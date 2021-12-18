package core_test

import (
	"github.com/charoleizer/my-collection/src/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Core", func() {

	It("should return 'Hello, 世界' ", func() {
		Expect(core.HelloSekai()).To(Equal("Hello, 世界"))
	})

})
