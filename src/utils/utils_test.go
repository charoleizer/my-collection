package utils_test

import (
	"github.com/charoleizer/my-collection/src/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	It("Should return 'Empty text' where there is no arguments", func() {
		Expect(utils.RemoveSpecialCharacters("")).To(Equal("Empty text"))
	})
})
