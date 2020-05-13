package goroutine

import (
	"fmt"
	"runtime"
)

// GetCPUNum ...
func GetCPUNum() {
	fmt.Println(runtime.NumCPU())
}
