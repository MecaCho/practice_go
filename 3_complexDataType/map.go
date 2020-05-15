package __complexDataType

import (
	"sync"
)

func MapHasNoLen() {

	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	// fmt.Println(m.Len())

}
