package __Interface

import "fmt"

func GetInterfaceType(x interface{}) string {

	switch x := x.(type) {

	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return x
	default:
		panic(fmt.Sprintf("unexpected type: %T: %v", x, x))
	}
}

func GetInt() int {
	a := [...]int{5: 3}
	fmt.Println(a)
	return 1
}

// func GetIntValueType() string {
// 	x := GetInt()
// 	switch x := x.(type) {
//
// 	case nil:
// 		return "NULL"
// 	case int, uint:
// 		return fmt.Sprintf("%d", x)
// 	case bool:
// 		if x {
// 			return "TRUE"
// 		}
// 		return "FALSE"
// 	case string:
// 		return x
// 	default:
// 		panic(fmt.Sprintf("unexpected type: %T: %v", x, x))
// 	}
// }
