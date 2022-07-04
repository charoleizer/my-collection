package domain_test

import (
	"testing"

	"github.com/charoleizer/my-collection/src/domain"
	"github.com/charoleizer/my-collection/src/models"
)

func TestFigure(t *testing.T) {
	assert := func(t *testing.T, result, expected interface{}) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%v', expected '%v'", result, expected)
		}
	}

	t.Run("Should return True if Figure's name is Luffy", func(t *testing.T) {
		result := domain.IsLuffy(models.Figures{"1", "Luffy"})
		expected := true
		assert(t, result, expected)
	})
}

func BenchmarkFigure(b *testing.B) {
	domain.IsLuffy(models.Figures{"1", "Luffy"})
}
