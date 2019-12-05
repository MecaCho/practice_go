package main

import (
	"strconv"
	"fmt"
)

//golang中字符串和各种int类型之间的相互转换方式：

func TestConvert()  {
	
	
	stringInput := "10"

	//string转成int：
	intRet, _ := strconv.Atoi(stringInput)
	fmt.Println(intRet)
	

	//string转成int64：
	int64Ret, _ := strconv.ParseInt(stringInput, 10, 64)
	fmt.Println(int64Ret)


	intInput := 10
	int64Input := int64(10)
	//int转成string：
	stringRet := strconv.Itoa(intInput)
	
	fmt.Println(stringRet)

	//int64转成string：
	stringRet = strconv.FormatInt(int64Input,10)

	fmt.Println(stringRet)


}

