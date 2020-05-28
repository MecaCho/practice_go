package gorotinueprint

import "fmt"

const (
	numNum   = 26
	numStr   = 26
	StrPrint = "ABCDEFGHIJKLMNOPQRSTUVWXYZAB"
)

func PrintChar() {

	for i := range StrPrint {
		fmt.Println(StrPrint[i])
	}

	s := []rune("people")

	for i, v := range s {
		fmt.Println(s[i], v)
	}

	s1 := "people"

	for i, v := range s1 {
		fmt.Println(s1[i], string(v))
	}
}

func PrintNumAndStr() {

	chanNum := make(chan bool)
	chanStr := make(chan bool)
	done := make(chan int)

	go func() {
		i := 0
		for {
			select {
			case <-chanNum:

				fmt.Print(i + 1)
				chanStr <- true
				if i >= numNum-1 {
					done <- 1
				}
				i++
				// default:
				// 	break

			}
		}
	}()

	go func() {

		i := 0
		for {
			select {
			case <-chanStr:
				fmt.Print(StrPrint[i : i+1])
				chanNum <- true
				i++
				// default:
				// 	break
			}
		}
	}()

	chanNum <- true
	<-done

}
