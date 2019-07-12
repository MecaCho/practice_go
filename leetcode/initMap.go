package leetcode

import "fmt"

type S struct {
	name string
}

func PrintMap() {
	m := map[string]S{"x": S{"one"}}
	m["X"] = S{"two"}

	fmt.Println("init map, can not assign map ... : ", m)
	//m["x"].name = "two"
}
