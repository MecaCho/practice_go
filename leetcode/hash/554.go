package hash

func leastBricks(wall [][]int) int {
	wallDict := map[int]int{}

	for _, wal := range wall {
		tmp := 0
		for _, w := range wal {
			tmp += w
			wallDict[tmp-1]++
		}
	}
	maxCount := 0
	for _, v := range wallDict {
		if maxCount > v {
			maxCount = v
		}
	}
	return len(wall) - maxCount
}

func LeastBricks(wall [][]int) int {
	return leastBricks(wall)
}

