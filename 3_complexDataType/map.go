package __complexDataType

import (
	"fmt"
	"sync"
)

func MapHasNoLen() {

	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	// fmt.Println(m.Len())

}

type Test struct {
	Name string
}

var list map[string]Test

func MapAddress() {
	list = make(map[string]Test)
	name := Test{"xiaoming"}
	list["name"] = name
	// can not assign
	// list["name"].Name = "Hello"
	fmt.Println(list["name"])
}
