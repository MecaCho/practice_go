package sort

import (
	"fmt"
	"sync"
	"testing"
)

func TestSort(t *testing.T) {
	Sort()

	var mux sync.Mutex
	mux.Lock()
	fmt.Println("test lock")
	mux.Unlock()
	mux.Unlock()
}
