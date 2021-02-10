package redis_go

import "testing"

func TestCounterWithRedisLock(t *testing.T) {
	CounterWithRedisLock()
}


// === RUN   TestCounterWithRedisLock
// Lock failed: false, <nil>
// Lock failed: false, <nil>
// Lock failed: false, <nil>
// Lock failed: false, <nil>
// Lock failed: false, <nil>
// Lock failed: false, <nil>
// Lock failed: false, <nil>
// Lock failed: false, <nil>
// Lock failed: false, <nil>
// Current counter value is: 1
// Unlock success.
// --- PASS: TestCounterWithRedisLock (0.01s)
// PASS