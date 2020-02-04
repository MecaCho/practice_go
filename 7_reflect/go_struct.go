package __reflect

import (
	"fmt"
	"reflect"
)

type T struct {
	N int
}

func ReStruct() {
	var n = T{42}
	fmt.Println(n.N)
	reflect.ValueOf(&n).Elem().FieldByName("N").SetInt(7)
	fmt.Println(n.N)
}
