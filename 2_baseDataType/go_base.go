package __baseDataType

import (
	"fmt"
	__channel "go_practice/9_channel"
	"go_practice/leetcode"
	"sync"
)

type Node struct {
	value int
	next  *Node
}

func init() {
	fmt.Println("hello world.")
}

func init() {
	fmt.Println("two init.")
}

func main() {
	SliceTest()
	leetcode.GetCode()
	leetcode.AddTest()
	leetcode.InitChannel()
	leetcode.PrintSlice()
	leetcode.PrintPnic()
	leetcode.PrintMap()
	leetcode.PrintSort()
	leetcode.Run()
	leetcode.RunPrint()
	leetcode.PrintAnonymousFunc()

	//leetcode.BufferChannel()
	//leetcode.WithoutBufferChannel()

	//channel.SendNilChannel()
	__channel.ReadCloseChannel()
	leetcode.InitSlice()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		leetcode.InitContext()
	}()
	wg.Wait()

	defer func() {
		fmt.Println("recover")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("ok")
	}()

	//
	DeferTest()

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

func DeferTest() {

	defer fmt.Println("abc")

	panic("test panic")

	defer fmt.Println("def")
}

func DeferPanicTest() {

	defer recover()

	panic("test panic")
}

func DeferPanicTest1() {

	defer func() {
		fmt.Println("recover......")
		recover()
	}()

	panic("test panic")
}

func mapTest() {
	fmt.Println("test map test, ")

	var hashmap map[string]string
	hashmap["qwq"] = "test_map"
	fmt.Println(hashmap["qwq"])

}

func proc() {
	panic("panic test.")

}

func SliceTest() {
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

func CapOfSlice() {

	s := []int{}

	fmt.Println("len: ", len(s), "   cap: ", cap(s))

	preCap := 0

	for i := 0; i < 1000; i++ {
		s = append(s, i)
		if cap(s) != preCap {
			fmt.Println("len: ", len(s), "   cap: ", cap(s))
			preCap = cap(s)
		}
	}
	// 	len:  1    cap:  1
	// len:  2    cap:  2
	// len:  3    cap:  4
	// len:  5    cap:  8
	// len:  9    cap:  16
	// len:  17    cap:  32
	// len:  33    cap:  64
	// len:  65    cap:  128
	// len:  129    cap:  256
	// len:  257    cap:  512
	// len:  513    cap:  1024
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

func Demo() {
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
