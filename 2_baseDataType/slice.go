package __baseDataType

import "fmt"

func SliceAppend() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

}

// new
func SliceAppend1() {

	// list := new([]int)
	// list = append(list, 1)
	// fmt.Println(list)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

func AddressConst() {
	const cl = 100
	var bl = 123

	// 	For an operand x of type T, the address operation &x generates a pointer of type *T to x.
	// The operand must be addressable, that is, either a variable, pointer indirection, or slice indexing operation;
	// or a field selector of an addressable struct operand;
	// or an array indexing operation of an addressable array.
	// As an exception to the addressability requirement, x may also be a (possibly parenthesized) composite literal.
	// If the evaluation of x would cause a run-time panic, then the evaluation of &x does too.
	println(&bl, bl)
	// println(&cl, cl)

	const k = 5
	v := k
	address := &v // This is allowed
	fmt.Println(address)
}

func reverse(a []int) {
	length := len(a)
	for i := 0; i < length/2; i++ {
		tmp := a[i]
		a[i] = a[length-i-1]
		a[length-i-1] = tmp
	}
}

func Reverse() {

	a := []int{1, 2, 3, 4, 5, 6}
	reverse(a)
	fmt.Println(a)
	// 	// Output: [3 2 1]

}
