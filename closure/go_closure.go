package closure

import "fmt"

func ClosureFunc(num int64)  {
	fmt.Println("closure outside: ", num)

	go func() {
		fmt.Println("closure func without params: ", num)

	}()

	go func(int64) {

		fmt.Println("closure func with params: ", num)
	}(num)


	go func(int64) {

		num = 20
		fmt.Println("closure func with params, num changed: : ", num)
	}(num)

	fmt.Println("num changed :", num)


}
