package utils_test

import (
	"github.com/charoleizer/my-collection/src/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	It("Should return empty string", func() {
		Expect(utils.RemoveSpecialCharacters("any text")).To(Equal(""))
	})
})
