package _2_sync

import (
	"fmt"
	"sync"
	"time"
)

var mu1 sync.RWMutex
var count int

// === RUN   TestRwMutex
// fatal error: all goroutines are asleep - deadlock!
func RwMutex() {
	go A1()
	time.Sleep(2 * time.Second)
	mu1.Lock()
	defer mu1.Unlock()
	count++
	fmt.Println(count)
}

// 15600272312
func A1() {
	mu1.RLock()
	defer mu1.RUnlock()
	B1()
}
func B1() {
	time.Sleep(5 * time.Second)
	C1()
}
func C1() {
	mu1.RLock()
	defer mu1.RUnlock()
}
