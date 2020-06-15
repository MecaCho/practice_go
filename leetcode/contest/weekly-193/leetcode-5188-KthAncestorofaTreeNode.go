package weekly_193

import "fmt"

func RangeString() {

	s := "fhjfhjdf在"

	for i, val := range s {
		fmt.Println(s[i])
		fmt.Println(i, val)
	}

	fmt.Println(len(s))

}
