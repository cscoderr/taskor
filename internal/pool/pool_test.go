package pool

import (
	"testing"
)

func TestNewWorkersLength(t *testing.T) {
	result := NewWorkers(10)
	expected := 10

	if len(result) != expected {
		t.Fatalf("expected %v, got %v", expected, len(result))
	}
}
