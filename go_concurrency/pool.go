package go_concurrency

import (
	"fmt"
	"sync"
	"time"
)

var bytePool = sync.Pool{
	New: newPool,
}

func newPool() interface{} {
	b := make([]byte, 1024)
	return &b

}

func SyncPool(n int) {
	s := time.Now().Unix()
	for i := 0; i < n; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
	e := time.Now().Unix()

	// fmt.Println("without pool: ", time.Since(s))

	e2 := time.Now().Unix()
	for j := 0; j < n; j++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
	e3 := time.Now().Unix()

	fmt.Println(e - s)
	fmt.Println(e3 - e2)
	// fmt.Println("with pool: ", time.Since(e2))
}
