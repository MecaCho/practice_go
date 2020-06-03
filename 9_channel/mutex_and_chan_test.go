package __channel

import "testing"

func BenchmarkRunWithChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunWithChan()
	}
}

// pkg: go_practice/9_channel
// 28773352	        41.3 ns/op
// PASS

func BenchmarkRunWithMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunWithMutex()
	}
}

// pkg: go_practice/9_channel
// 106226368	        10.3 ns/op
// PASS

func TestTwoPrint(t *testing.T) {
	TwoPrint()
}
