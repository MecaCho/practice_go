package main

import (
	"fmt"
)

var x complex128 = complex(1,2)

var y complex128 = complex(3, 4)

const (
	_ = uint64(1) << (10 * iota)
	KiB //1024
	MiB //1048576
	GiB //1073741824
	TiB //1099511627776  超过int32
	PiB //1125899906842624
	EiB //1152921504606846976
	ZiB
	YiB
)

func main()  {

	fmt.Println(x*y)
	fmt.Println(real(x*y))
	fmt.Println(imag(x*y))

	// const IPv4Len = 4

	// var p [IPv4Len]byte

	fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB)

}