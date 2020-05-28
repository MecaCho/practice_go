package _2_sync

import (
	"fmt"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	WaitGroup()
	// 	panic: sync: WaitGroup is reused before previous Wait has returned [recovered]
	// 	panic: sync: WaitGroup is reused before previous Wait has returned
	//
	// goroutine 18 [running]:
	// testing.tRunner.func1(0xc0000de100)
	// 	/usr/local/go/src/testing/testing.go:874 +0x3a3
	// panic(0x1111960, 0x116c960)
	// 	/usr/local/go/src/runtime/panic.go:679 +0x1b2
	// sync.(*WaitGroup).Wait(0xc00009a1d0)
	// 	/usr/local/go/src/sync/waitgroup.go:132 +0xad
	// go_practice/12_sync.WaitGroup()
	// 	/Users/rainmc/GO/src/go_practice/12_sync/waitgroup.go:22 +0x79
	// go_practice/12_sync.TestWaitGroup(0xc0000de100)
	// 	/Users/rainmc/GO/src/go_practice/12_sync/waitgroup_test.go:6 +0x20
	// testing.tRunner(0xc0000de100, 0x1150370)
	// 	/usr/local/go/src/testing/testing.go:909 +0xc9
	// created by testing.(*T).Run
	// 	/usr/local/go/src/testing/testing.go:960 +0x350
}

func TestDeferDone(t *testing.T) {
	DeferDone()
}

func TestPrintRandIntMap(t *testing.T) {
	PrintRandIntMap()
}

func TestWaitGroupWithMaxProcs(t *testing.T) {
	WaitGroupWithMaxProcs()
}

func TestIsPrefixOfWord(t *testing.T) {
	res := IsPrefixOfWord("i use triple pillow", "pill")
	fmt.Println(res)
}

func TestVerifyPeople(t *testing.T) {
	res := VerifyPeople("peabcoplbe")
	fmt.Println(res)

	res = VerifyPeople("aapeoplea")
	fmt.Println(res)

	res = VerifyPeople("peaaple")
	fmt.Println(res)
}
