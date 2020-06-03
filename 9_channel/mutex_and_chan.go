package __channel

import (
	"sync"
)

var mutex sync.Mutex

func RunWithMutex() {
	mutex.Lock()
	Woker()
	mutex.Unlock()
}

var ch = make(chan int, 1)

func RunWithChan() {
	ch <- 1
	Woker()
	<-ch
}

func Woker() {
	// mapTest := make(map[string]string, 10)
}

var a string
var once sync.Once

func setup() {
	a = "hello, world"
}
func doprint() {
	once.Do(setup)
	print(a)
}

func TwoPrint() {
	go doprint()
	go doprint()
}
