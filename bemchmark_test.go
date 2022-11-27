package main

import "testing"

func BenchmarkA(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solveA(10000)
	}
}
