package greedy

import (
	"fmt"
	"sort"
)

func MaxEvents(events [][]int) int {

	sort.Slice(events, func(i, j int) bool {

		return events[i][1] < events[j][1]
	})

	fmt.Println(events)

	attends := [100000]int{0}

	for _, event := range events {
		for i := event[0]; i <= event[1]; i++ {
			if attends[i] == 0 {
				attends[i] = 1
				break
			}
		}
	}

	res := 0

	for k := range attends {
		res += attends[k]
	}

	return res

}
