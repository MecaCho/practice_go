package main

import "fmt"

type Node struct {
	value int
	next  *Node
}



func main(){
	//
	sliceTest()

	defer func() {
		fmt.Println("recover")
		if err := recover();err != nil {
			fmt.Println(err)
		}
		fmt.Println("ok")
	}()

	//
	deferTest()

	defer fmt.Println("123")
	defer fmt.Println("456")
	nodeList := make([]Node, 10)
	for i := 0; i < 9; i++ {
	nodeList[i].value = i

	nodeList[i].next = &nodeList[i+1]
	}
	nodeList[9].value = 9
	//var node Node
	for i, node := range nodeList {
	fmt.Println(i, node.value)
	}
	//reverseList(*nodeList)
//node.next

}

func deferTest()  {

	defer fmt.Println("abc")

	panic("test panic")

	defer fmt.Println("def")

}

func mapTest()  {

	hashmap := map[]string
	

}


func proc()  {
	panic("panic test.")

}

func sliceTest()  {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

}

func reverseList(head *Node) *Node {
	var tmp *Node = nil
	for head != nil {
		tmp = head.next
		head.next = tmp.next
		tmp.next = head.next
	}
	return tmp

}
