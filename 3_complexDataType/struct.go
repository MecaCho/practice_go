package __complexDataType

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

// 这是Golang的组合模式，可以实现OOP的继承。
// 被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法（一定要是匿名字段），
// 但它们的方法(ShowA())调用时接受者并没有发生变化。
// 此时People类型并不知道自己会被什么类型组合，当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。
func StructCom() {
	t := Teacher{}
	t.ShowA()
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

// ambiguous selector my.m1
func TypeDefine3() {
	my := MyStruct{}
	// my.m1()
	my.T1.m1()
	// my.m1()

}

func StructEqual() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	// sn3 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}
	// sn4 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}

	// // sn3与sn1就不是相同的结构体了，不能比较。
	// // 还有一点需要注意的是结构体是相同的，但是结构体属性中有不可以比较的类型，如map,slice。
	// // 如果该结构属性都是可以比较的，那么就可以使用“==”进行比较操作。
	// if sn3 == sn4 {
	// 	fmt.Println("sn3 == sn4")
	// }
}
