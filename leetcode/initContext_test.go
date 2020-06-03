package leetcode

import "testing"

func TestInitContext(t *testing.T) {
	InitContext()
}

func TestInitContext1(t *testing.T) {
	InitContext1()
}

func BenchmarkInitContext1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InitContext1()
	}
}

func TestFallThrough(t *testing.T) {
	FallThrough()
}
