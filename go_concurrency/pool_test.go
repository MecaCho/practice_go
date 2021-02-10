package go_concurrency

import (
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	SyncPool(100000000)
	// SyncPool(1000000000)
}

var bytepool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func BenchmarkWithoutSyncPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
}

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkWithoutSyncPool-16    	1000000000	         0.260 ns/op
// PASS

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkWithoutSyncPool-16    	1000000000	         0.530 ns/op
// PASS

func BenchmarkWithSyncPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
}

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkWithSyncPool-16    	77769235	        14.0 ns/op
// PASS

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkWithSyncPool-16    	 3038128	       369 ns/op
// PASS

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkSyncPool-16    	1000000000	         0.260 ns/op
// PASS

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkSyncPool2-16    	 3186481	       377 ns/op
// PASS

type Node struct {
	Name string
}

var nodePool = sync.Pool{
	New: func() interface{} {
		return new(Node)
	},
}

func BenchmarkWithoutSyncPool1(b *testing.B) {
	var node *Node
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// for j := 0; j < 10000; j ++{
		// 	node = new(Node)
		// 	node.Name = "test"
		// }
		node = new(Node)
		node.Name = "test"
	}
}

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkWithoutSyncPool1-16    	1000000000	         0.509 ns/op	       0 B/op	       0 allocs/op
// PASS

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkWithoutSyncPool1-16    	  213987	      5178 ns/op	       0 B/op	       0 allocs/op
// PASS

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkWithoutSyncPool1-16    	    4324	    246215 ns/op	  160001 B/op	   10000 allocs/op
// PASS

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkWithoutSyncPool1-16    	52674117	        23.0 ns/op	      16 B/op	       1 allocs/op
// PASS

func BenchmarkWithSyncPool1(b *testing.B) {
	var node *Node
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// for j := 0; j < 10000; j ++{
		// 	node = nodePool.Get().(*Node)
		// 	node.Name = "test"
		// 	nodePool.Put(node)
		// }
		node = nodePool.Get().(*Node)
		node.Name = "test"
		nodePool.Put(node)
	}
}

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkWithSyncPool1-16    	    6549	    163542 ns/op	       0 B/op	       0 allocs/op
// PASS

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkWithSyncPool1-16    	    8138	    157292 ns/op	       0 B/op	       0 allocs/op
// PASS

// goos: darwin
// goarch: amd64
// pkg: go_practice/go_concurrency
// BenchmarkWithSyncPool1-16    	71456332	        14.8 ns/op	       0 B/op	       0 allocs/op
// PASS
