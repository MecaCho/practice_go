package redis_go

import "testing"

func TestCounterInr(t *testing.T) {
	for i := 0; i < 10; i++{
		CounterInr()
	}
}


// === RUN   TestCounterInr
// 910
// 863
// 912
// 884
// 865
// 865
// 882
// 800
// 836
// 864
// --- PASS: TestCounterInr (0.01s)
// PASS
