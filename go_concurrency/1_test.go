package go_concurrency

import "testing"

func TestCount(t *testing.T) {
	for i := 0; i < 3; i++ {
		Count()
	}
	// 	310909
	// 224339
	// 198934
}
