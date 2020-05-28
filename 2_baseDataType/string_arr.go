package __baseDataType

import "fmt"

func StringArr() {
	//字符串使用
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	for arr := range array2 {
		fmt.Println(arr)
	}
}

func StringIndex() {

	s := "abcdefg"
	// s[1] = "1"

	sBytes := []byte(s)
	sBytes[0] = []byte("Q")[0]
	fmt.Println(sBytes, string(sBytes))

	sRunes := []rune(s)
	sRunes[0] = []rune("在")[0]
	fmt.Println(sRunes, string(sRunes))
	// 	=== RUN   TestStringIndex
	// [81 98 99 100 101 102 103] Qbcdefg
	// [22312 98 99 100 101 102 103] 在bcdefg
}

func StringIndex1() {

	s := "abcdef"
	fmt.Println(s[0])
	fmt.Println()

}
