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

func DeferDone() {
	const N = 10

	var wg = &sync.WaitGroup{}

	for i := 0; i < N; i++ {
		go func(i int) {
			// defer wg.Done()
			wg.Add(1)
			println(i)
			defer wg.Done()
			// wg.Done()
		}(i)
	}
	wg.Wait()
}

// panic: sync: WaitGroup is reused before previous Wait has returned [recovered]
// panic: sync: WaitGroup is reused before previous Wait has returned
//
// goroutine 33 [running]:
// testing.tRunner.func1(0xc0000e0100)
// /usr/local/go/src/testing/testing.go:874 +0x3a3
// panic(0x1111960, 0x116c960)
// /usr/local/go/src/runtime/panic.go:679 +0x1b2
// sync.(*WaitGroup).Wait(0xc00009a1d0)
// /usr/local/go/src/sync/waitgroup.go:132 +0xad
// go_practice/12_sync.DeferDone()
// /Users/rainmc/GO/src/go_practice/12_sync/waitgroup.go:40 +0x7f
// go_practice/12_sync.TestDeferDone(0xc0000e0100)
// /Users/rainmc/GO/src/go_practice/12_sync/waitgroup_test.go:28 +0x20
// testing.tRunner(0xc0000e0100, 0x1150358)
// /usr/local/go/src/testing/testing.go:909 +0xc9
// created by testing.(*T).Run
// /usr/local/go/src/testing/testing.go:960 +0x350
//
// Process finished with exit code 1

// 9
//
// In your recursive cases in ThreadSafeCalcWords, you're calling wg.Done before calling wg.Add. That means that the wg can drop down to 0 (which will trigger the Wait to complete) before you actually finish all the work. Calling Add again while the Wait is still in the process of resolving is what triggers the error, but more importantly, it probably just plain isn't what you want.
//
// Change the order of operations so that you always Add any new work before doing Done on the existing work, and the Wait won't trigger prematurely. The simplest way to accomplish this would probably be a single call to wg.Done() at the bottom of the function, or a single defer at the top, and removing all the others.
