package __complexDataType

import (
	"fmt"
	"sync"
)

func MapHasNoLen() {

	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	// fmt.Println(m.Len())

}

type Test struct {
	Name string
}

var list map[string]Test

func MapAddress() {
	list = make(map[string]Test)
	name := Test{"xiaoming"}
	list["name"] = name
	// can not assign
	// list["name"].Name = "Hello"
	fmt.Println(list["name"])
}

type student struct {
	Name string
	Age  int
}

func PaserStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	// for k, stu := range stus {
	// 	m[stu.Name] = &stus[k]
	// }

	for k, v := range m {
		fmt.Println(k, v)
	}
	// fmt.Println(m)
}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func ConMap() {
	m := sync.Mutex{}
	ages := make(map[string]int, 10)
	ua := UserAges{ages, m}
	for i := 0; i < 10; i++ {
		go func(i int) {
			name := fmt.Sprintf("name-%d", i)
			ua.Add(name, i)
			age := ua.Get(name)
			fmt.Println(age)
		}(i)
		// ua.Add(fmt.Sprintf("name-%d", i), i)
	}
	for k, v := range ua.ages {
		fmt.Println(k, v)
	}

}

func GetValue(m map[int]string, id int) (string, bool) {
	if v, exist := m[id]; exist {
		return v, true
	}
	return "", false
}
func GetMapValue() {
	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}

	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)
}

func SyncMap() {

	newMap := sync.Map{}
	newMap.Store("abc", "123")
	newMap.LoadOrStore("qwq", "qwe")
	newMap.Load("qwq")
	newMap.Delete("qwq")

}
