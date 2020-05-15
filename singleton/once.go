package singleton

import (
	"fmt"
	"sync"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

func Task() {
	fmt.Println("do task")
}

func OnceRun() {
	m := sync.Mutex{}
	done := uint32(0)
	o := Once{m, done}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			o.Do(Task)
			wg.Done()
		}()
	}

	wg.Wait()

}
