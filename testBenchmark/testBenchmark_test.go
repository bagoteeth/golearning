package testBenchmark

import (
	"golearning/testFunc"
	"testing"
)

func BenchmarkDo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testFunc.TryBenchmark()
	}
}
func BenchmarkFast(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testFunc.TryFastBenchmark()
	}
}
