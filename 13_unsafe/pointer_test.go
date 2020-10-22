package _3_unsafe

import (
	"fmt"
	"reflect"
	"strings"
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
	res := strings.Contains("23234,789", "123")
	fmt.Println(res)

	res = strings.Contains("23234,789", "789")
	fmt.Println(res)
}

func TestBytesToString2(t *testing.T) {
	s := []int{1,2,3,4,5,6}
	res := []int{}
	for _, v := range s{
		res = append(res, v)
	}
	fmt.Println(s)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
	fmt.Println(res)
}

