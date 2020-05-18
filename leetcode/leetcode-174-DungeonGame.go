package leetcode

func calculateMinimumHP(dungeon [][]int) int {

	colLen := len(dungeon)
	rowLen := len(dungeon[0])
	if colLen+rowLen < 2 {
		return 0
	}

	for i := 0; i < colLen; i++ {
		for j := 0; j < rowLen; j++ {
			if i == 0 && j != 0 {
				dungeon[i][j] += dungeon[i][j-1]
			} else if j == 0 && i != 0 {
				dungeon[i][j] += dungeon[i-1][j]
			} else if i != 0 && j != 0 {
				if dungeon[i][j-1] < dungeon[i-1][j] {
					dungeon[i][j] += dungeon[i][j-1]
				} else {
					dungeon[i][j] += dungeon[i-1][j]
				}
			}

		}
	}
	return dungeon[colLen-1][rowLen-1]

}
