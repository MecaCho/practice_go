package main

import (
	"fmt"
	"reflect"
)

func main1() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
}
