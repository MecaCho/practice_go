package __reflect

import (
	"reflect"
	"fmt"
)

func IsEqual()  {

	var map1, map2 map[string]int = nil, make(map[string]int)
	res := reflect.DeepEqual(map1, map2)
	fmt.Println(res)

	var slice1, slice2 []int = nil, []int{}
	res1 := reflect.DeepEqual(slice1, slice2)
	fmt.Println(res1)
}
