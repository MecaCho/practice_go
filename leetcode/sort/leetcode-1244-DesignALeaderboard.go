package sort

import "sort"

type Leaderboard struct {
	scores map[int]int
}


func Constructor() Leaderboard {
	scores := make(map[int]int, 1)
	b := Leaderboard{scores}
	return  b
}


func (this *Leaderboard) AddScore(playerId int, score int)  {
	if _, ok := this.scores[playerId]; ok{
		this.scores[playerId] += score
	}else {
		this.scores[playerId] = score
	}
}


func (this *Leaderboard) Top(K int) int {

	allScores := []int{}

	for k, _ := range this.scores{
		allScores = append(allScores, this.scores[k])
	}

	sort.Ints(allScores)

	res := 0
	for i := 0; i < K; i++{
		res += allScores[len(allScores) - 1 - i]
	}

	return res
}


func (this *Leaderboard) Reset(playerId int)  {
	this.scores[playerId] = 0
}


/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */
