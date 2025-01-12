package __baseDataType

import (
	"fmt"
	"unicode/utf8"
)

func RuneBasic() {
	// golang中string底层是通过byte数组实现的，len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	stringTest := "abc123中文"
	fmt.Println(len(stringTest)) // 12
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	fmt.Println(len(stringTest), len([]rune(stringTest)), utf8.RuneCountInString(stringTest)) // 12 8 8
}

func GetStringLen() {
	s := "abc123中文"

	fmt.Println(len(s))

	fmt.Println(len([]rune(s)))

}
