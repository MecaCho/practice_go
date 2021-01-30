package __Interface

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
)

type S struct {
}

func f(x interface{}) {
	var w io.Writer
	fmt.Println(w)
	w = os.Stdout
	w = new(bytes.Buffer)
	// aa := &bytes.Buffer{}
	// aa.
	runtime.GOMAXPROCS(1)
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

// var w io.Writer = os.Stdout

// ff, ok := w.(*os.File)
// bb, ok := w.(*bytes.Buffer)
