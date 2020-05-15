package _2_sync

import (
	"sync"
	"time"
)

// === RUN   TestWaitGroup
// --- FAIL: TestWaitGroup (0.00s)
// panic: sync: WaitGroup is reused before previous Wait has returned [recovered]
// 	panic: sync: WaitGroup is reused before previous Wait has returned
func WaitGroup() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		// panic
		wg.Add(1)
	}()
	wg.Wait()

}
