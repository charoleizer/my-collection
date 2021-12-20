package utils_test

import (
	"github.com/charoleizer/my-collection/src/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {

	It("Should return 'Empty text' when there is no arguments", func() {
		Expect(utils.RemoveSpecialCharacters("")).To(Equal("Empty text"))
	})

	It("Should return 'Expected text' when argument is '[-!,Expected*)@#%( text&$_?.^' ", func() {
		Expect(utils.RemoveSpecialCharacters("[-!,Expected*)@#%( text&$?.^")).To(Equal("Expected text"))
	})

})
