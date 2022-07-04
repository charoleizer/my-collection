package domain_test

import (
	"testing"

	"github.com/charoleizer/my-collection/src/domain"
)

func TestHello(t *testing.T) {
	assert := func(t *testing.T, result, expected interface{}) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%v', expected '%v'", result, expected)
		}
	}

	t.Run("Should return 'Hello, 世界'", func(t *testing.T) {
		result := domain.HelloSekai()
		expected := "Hello, 世界"
		assert(t, result, expected)
	})
}

func BenchmarkHello(b *testing.B) {
	domain.HelloSekai()
}
