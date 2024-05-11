package log

import (
	"testing"
)

func BenchmarkError(b *testing.B) {
	// run the Debug function b.N times
	for i := 0; i < b.N; i++ {
		Info("", "%s %s", "param1", "param2")
		Error("", "%s %s", "param1", "param2")
	}
}
