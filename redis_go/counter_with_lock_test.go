package redis_go

import "testing"

func TestCounterInrWithLock(t *testing.T) {
	for i := 0; i < 10; i++{
		CounterInrWithLock()
	}
}

// === RUN   TestCounterInrWithLock
// 1000
// 1000
// 1000
// 1000
// 1000
// 1000
// 1000
// 1000
// 1000
// 1000
// --- PASS: TestCounterInrWithLock (0.01s)
// PASS
