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

	//
	mapTest()

	//

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
	fmt.Println("test map test, ")

	var hashmap map[string]string
	hashmap["qwq"] = "test_map"
	fmt.Println(hashmap["qwq"])


}


func proc()  {
	panic("panic test.")

}

func sliceTest()  {
	s := make([]int, 5)
	fmt.Printf("slice len: %d, cap: %d\n", len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Printf("slice len: %d, cap: %d", len(s), cap(s))
	fmt.Println(s)

	s = make([]int, 5, 6)
	fmt.Printf("slice len: %d, cap: %d\n", len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Printf("slice len: %d, cap: %d\n", len(s), cap(s))

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

func Demo(){
	var sli []string
	abc := append(sli, "abc")
	fmt.Println(abc)

	var i *int

	i = new(int)

	*i = 10
	fmt.Println(*i)

	type Entity struct {
		Desc string `json:"desc"`
	}

	testMap := make(map[string]Entity, 0)

	testMap["qwq"] = Entity{Desc: "This is a test"}

	fmt.Println(testMap["qwq"].Desc)

	var testMap1 map[string]Entity

	testMap1["qwq"] = Entity{Desc: "test"}
}
