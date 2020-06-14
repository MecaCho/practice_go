package __base

import (
	"fmt"
	"unsafe"
)

func SizeOf() {
	i := int(1)
	j := int64(1)
	fmt.Println(unsafe.Sizeof(i), unsafe.Sizeof(j))
}
