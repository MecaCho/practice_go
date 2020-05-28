package gorotinueprint

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrintNumAndStr(t *testing.T) {
	PrintNumAndStr()
	strings.Index(StrPrint, "A")
	fmt.Println(StrPrint[0:1])
}

func TestPrintChar(t *testing.T) {
	PrintChar()
}
