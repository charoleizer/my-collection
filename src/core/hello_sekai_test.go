package core_test

import (
	"testing"

	"github.com/charoleizer/my-collection/src/core"
)

func TestHelloSekai(t *testing.T) {
	result := core.HelloSekai()
	if result != "Hello, 世界" {
		t.Error("result should be 'Hello, 世界', got", result)
	}
}
