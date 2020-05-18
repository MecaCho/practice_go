package __baseDataType

import "fmt"

type User struct {
}

type MyUser1 User
type MyUser2 = User

func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}

func (i User) m2() {
	fmt.Println("User.m2")
}

func TypeDefine1() {
	var i1 MyUser1
	var i2 MyUser2
	i1.m1()
	i2.m2()
}

type T1 struct {
}

func (t T1) m1() {
	fmt.Println("T1.m1")
}

type T2 = T1

type MyStruct struct {
	T1
	T2
}

func TypeDefine3() {
	my := MyStruct{}
	// my.m1()
	my.T1.m1()

}

func TypeDefine() {
	type MyInt1 int
	type MyInt2 = int
	var i int = 9
	// cannot use i (type int) as type MyInt1 in assignment
	// var i1 MyInt1 = i
	var i2 MyInt2 = i
	// fmt.Println(i1, i2)
	fmt.Println(i2)
}
