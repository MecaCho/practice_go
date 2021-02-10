package slide_window

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	res := CheckInclusion("abc", "afhjhfjksabc")
	fmt.Println(res)

	// 单引号字符， rune
	fmt.Println('c' - 'a')

	// 双引号字符串
	// fmt.Println("a" - "a")

	// ``可以带换行符
	// fmt.Println(`a` - `a`)
}
