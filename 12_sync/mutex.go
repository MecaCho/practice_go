package _2_sync

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var chain string

func A() {
	mu.Lock()
	defer mu.Lock()
	chain = chain + " --> A"
	B()
}
func B() {
	chain = chain + " --> B"
	C()
}
func C() {
	mu.Lock()
	defer mu.Lock()
	chain = chain + " --> C"
}

// === RUN   TestMutexLock
// fatal error: all goroutines are asleep - deadlock!

func MutexLock() {
	chain = "main"
	A()
	fmt.Println(chain)
}
