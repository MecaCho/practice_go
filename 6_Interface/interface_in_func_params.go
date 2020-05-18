package __Interface

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func FuncParams() {
	s := S{}
	p := &s
	f(s)
	// g(s)
	f(p)
	// g(p)
	// 	./interface_in_func_params.go:16:3: cannot use s (type S) as type *interface {} in argument to g:
	// 	*interface {} is pointer to interface, not interface
	// ./interface_in_func_params.go:18:3: cannot use p (type *S) as type *interface {} in argument to g:
	// 	*interface {} is pointer to interface, not interface
}
