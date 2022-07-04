package utils_test

import (
	"testing"

	"github.com/charoleizer/my-collection/src/utils"
)

func TestHello(t *testing.T) {
	assert := func(t *testing.T, result, expected interface{}) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%v', expected '%v'", result, expected)
		}
	}

	t.Run("Should return 'Empty text' when there is no arguments'", func(t *testing.T) {
		result := utils.RemoveSpecialCharacters("")
		expected := "Empty text"
		assert(t, result, expected)
	})

	t.Run("Should return 'Expected text' when argument is '[-!,Expected*)@#%( text&$_?.^' ", func(t *testing.T) {
		result := utils.RemoveSpecialCharacters("[-!,Expected*)@#%( text&$?.^")
		expected := "Expected text"
		assert(t, result, expected)
	})
}
