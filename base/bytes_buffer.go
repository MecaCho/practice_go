package base

import (
	"bytes"
	"fmt"
	"os"
)

// 创建缓冲器
func CreateBuffer() {

	buff1 := bytes.NewBufferString("hello")
	fmt.Println(buff1.String())

	buff2 := bytes.NewBuffer([]byte("hello 1"))
	fmt.Println(buff2.String())
}

// 写入缓冲器
func BufferWrite() {
	bufferNew := bytes.NewBufferString("hello")

	bufferNew.Write([]byte(" +world."))

	fmt.Println(bufferNew.String())

	bufferNew.WriteString(" +world..")
	fmt.Println(bufferNew.String())

	sRune := " +中文rune"
	// var s rune
	// s = "中文rune"
	// const r = '☺'
	for k, r := range sRune {
		fmt.Println(k, r)
		bufferNew.WriteRune(r)
	}
	fmt.Println(bufferNew)
}

func BufferRead() {

}

// 从缓冲器写出到文件
func BufferWriteToFile() (err error) {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	buff := bytes.NewBufferString("hello world")
	buff.WriteTo(file)
	// fmt.Fprint(file, buff.String())
	return
}

func BufferReadFromFile() {

}
