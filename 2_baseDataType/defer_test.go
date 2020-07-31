package __baseDataType

import (
	"fmt"
	"testing"
)

func TestDeferReturn(t *testing.T) {
	DeferReturn()
}

func TestDeferstack(t *testing.T) {
	Deferstack()
}

func TestCalDefer(t *testing.T) {
	CalDefer()
}

func TestDeferFun(t *testing.T) {
	DeferFun()
}

func TestDeferFunc1(t *testing.T) {
	res := DeferFunc1(1)
	fmt.Println(res)
}

func TestVarInDefer(t *testing.T) {
	VarInDefer()
	VarInDefer1()
}

func TestDeferstack2(t *testing.T) {
	Deferstack()
}
