package _3_unsafe

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestBytesToString(t *testing.T) {
	s := "abcdefg"
	res := StringToBytes(s)
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	fmt.Printf("%+v, %+v\n", &s, &stringHeader.Data)
	fmt.Println(res)

	bytes := []byte(s)
	res1 := BytesToString(bytes)
	fmt.Println(res1)
}

func TestStringToBytes(t *testing.T) {
	c := 10
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&c)) + uintptr(1))))
}
