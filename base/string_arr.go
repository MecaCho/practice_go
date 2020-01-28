package base

import "fmt"

func StringArr() {
	//字符串使用
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	for arr := range array2 {
		fmt.Println(arr)
	}
}
