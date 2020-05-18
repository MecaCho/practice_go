package __baseDataType

import "fmt"

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

func GetConst() {
	fmt.Println(x, y, z, k, p)
}
