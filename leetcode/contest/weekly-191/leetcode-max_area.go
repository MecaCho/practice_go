package weekly_191

import "math"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = append(horizontalCuts, 0, h)

	verticalCuts = append(verticalCuts, 0, w)

	max_h := 0
	hs := []int{}
	for i, v := range horizontalCuts {
		if i > 0 {
			h := v - horizontalCuts[i-1]
			if h > max_h {
				max_h = h
			}
			hs = append(hs, h)
		}
	}
	max_v := 0
	vs := []int{}
	for i, val := range verticalCuts {
		if i > 0 {
			v := val - verticalCuts[i-1]
			if v > max_v {
				max_v = v
			}
			vs = append(vs, v)
		}
	}

	return (max_h * max_v) % (int(math.Pow10(9)) + 7)

}
