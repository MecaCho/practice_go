package __function

import (
	"fmt"
	"log"
	"strings"
)

func Add1(r rune) rune {
	return r + 1
}

func main() {
	log.SetPrefix("xxxx: ")

	log.SetFlags(0)
	log.Println("test log;")

	fmt.Println(strings.Map(Add1, "test"))

	fmt.Println(strings.Map(func(x rune) rune { return x + 1 }, "1234"))
}
