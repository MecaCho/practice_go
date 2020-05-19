package __baseDataType

import (
	"fmt"
)

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

const (
	a int64 = 1024
	b       = 123
)

const Pi float64 = 3.14

const zero = 0.0

const pp, qq, rr int = 1, 2, 3

const xx, yy, zz = 1, 2, 3

// const ErrorN  = errors.New("fjshfkj")

func GetConst() {
	fmt.Println(x, y, z, k, p)
}
