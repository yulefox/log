package log

import (
	"strings"
	"testing"
)

func TestGetStacks(t *testing.T) {
	stacks := GetStacks(5)
	for _, stack := range stacks {
		Info("", stack)
	}

	// Check if GetStacks returns a non-empty slice
	if len(stacks) == 0 {
		t.Errorf("GetStacks returned an empty slice")
	}

	// Check if the first element of the slice contains the expected string
	expected := "runtime/debug.Stack"
	if !strings.Contains(stacks[0], expected) {
		t.Errorf("Expected the first element of the slice to contain %s, got %s", expected, stacks[0])
	}
}
