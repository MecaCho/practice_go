package __base

import "fmt"

func Point()  {

	x := 1
	y := &x


	fmt.Println(x, *y, y)
	x = 2
	// *y = 2

	fmt.Println(x, *y, y)

}
