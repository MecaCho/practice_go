package redis_go

import (
	"fmt"
	"sync"
)



func CounterInrWithLock() {
	var counter int
	var wg sync.WaitGroup
	var l sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			defer l.Unlock()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
