package __reflect

import (
	"fmt"
	"reflect"
)

func GetType() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)
	// 	=== RUN   TestGetType
	// int
	// int
}
