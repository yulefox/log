package log

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetStack(t *testing.T) {
	stack := GetStack(2, 10)
	fmt.Println(strings.Join(toStrings(stack), "\n"))

	// Check if the first element of the slice contains the expected string
	expected := "TestGetStack"
	if !strings.Contains(toString(stack[0]), expected) {
		t.Errorf("Expected the first element of the slice to contain %s, got %v", expected, stack)
	}
}
