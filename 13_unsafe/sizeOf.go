package _3_unsafe

import (
	"fmt"
	"unsafe"
)

type Struct1 struct {
	c bool
	a float64
	b int16
}

type Struct2 struct {
	a float64
	b int16
	c bool
}

type Struct3 struct {
	c bool
	b int16
	a float64
}

var boolTest bool

var stringTest string

var mapTest map[int]int

var funcTest  func()

var chanTest chan int

var interfaceTest interface{}


func SizeOf()  {

	fmt.Println("Size of float64: ", unsafe.Sizeof(float64(0)))

	fmt.Println("Size of bool: ", unsafe.Sizeof(boolTest))

	fmt.Println("Size of string: ", unsafe.Sizeof(stringTest))

	fmt.Println("Size of map: ", unsafe.Sizeof(mapTest))

	fmt.Println("Size of func: ", unsafe.Sizeof(funcTest))

	fmt.Println("Size of channel: ", unsafe.Sizeof(chanTest))

	fmt.Println("Size of interface: ", unsafe.Sizeof(interfaceTest))

	s1 := Struct1{}
	fmt.Println("Size of struct1: ", unsafe.Sizeof(Struct1{}),
		"\n Alignof struct1: ", unsafe.Alignof(Struct1{}),
		"\n\n Alignof struct1.a: ", unsafe.Alignof(s1.a),
		"\n Alignof struct1.b: ", unsafe.Alignof(s1.b),
		"\n Alignof struct1.c: ", unsafe.Alignof(s1.c),
		"\n\n Offsetof struct1.a: ", unsafe.Offsetof(s1.a),
		"\n Offsetof struct1.b: ", unsafe.Offsetof(s1.b),
		"\n Offsetof struct1.c: ", unsafe.Offsetof(s1.c))

	fmt.Println("Size of struct2: ", unsafe.Sizeof(Struct2{}),
		"\n Alignof struct1: ", unsafe.Alignof(Struct1{}),
		"\n\n Alignof struct1.a: ", unsafe.Alignof(s1.a),
		"\n Alignof struct1.b: ", unsafe.Alignof(s1.b),
		"\n Alignof struct1.c: ", unsafe.Alignof(s1.c),
		"\n\n Offsetof struct1.a: ", unsafe.Offsetof(s1.a),
		"\n Offsetof struct1.b: ", unsafe.Offsetof(s1.b),
		"\n Offsetof struct1.c: ", unsafe.Offsetof(s1.c))

	fmt.Println("Size of struct3: ", unsafe.Sizeof(Struct3{}),
		"\n Alignof struct1: ", unsafe.Alignof(Struct1{}),
		"\n\n Alignof struct1.a: ", unsafe.Alignof(s1.a),
		"\n Alignof struct1.b: ", unsafe.Alignof(s1.b),
		"\n Alignof struct1.c: ", unsafe.Alignof(s1.c),
		"\n\n Offsetof struct1.a: ", unsafe.Offsetof(s1.a),
		"\n Offsetof struct1.b: ", unsafe.Offsetof(s1.b),
		"\n Offsetof struct1.c: ", unsafe.Offsetof(s1.c))
}
