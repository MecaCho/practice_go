package __baseDataType

import "fmt"

func SliceAppend() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

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
