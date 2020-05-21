package leetcode

func countCornerRectangles(grid [][]int) int {
	res := 0
	count := make(map[[2]int]int)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				for k := j + 1; k < len(grid[i]); k++ {
					if grid[i][k] == 1 {

						if v, ok := count[[2]int{j, k}]; ok {
							res += v
							count[[2]int{j, k}] += 1
						} else {
							count[[2]int{j, k}] = 1
						}

					}
				}
			}
		}
	}

	return res
}
