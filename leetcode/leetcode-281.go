package leetcode

type ZigzagIterator struct {
	Vals []int
	Num  int
	Loc  int
}

func Constructor(v1, v2 []int) *ZigzagIterator {

	max_num := len(v2)
	if len(v1) > len(v2) {
		max_num = len(v1)
	}
	z := ZigzagIterator{}
	z.Num = len(v1) + len(v2)
	z.Loc = 0
	for i := 0; i < max_num; i++ {
		if i < len(v1) {
			z.Vals = append(z.Vals, v1[i])
		}

		if i < len(v2) {
			z.Vals = append(z.Vals, v2[i])
		}

	}
	return &z

}

func (this *ZigzagIterator) next() int {
	val := this.Vals[this.Loc]
	this.Loc++

	return val

}

func (this *ZigzagIterator) hasNext() bool {
	res := this.Loc < this.Num
	// fmt.Println(res)
	return res
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */
