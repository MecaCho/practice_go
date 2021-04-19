package slide_window



func MaxSatisfied(customers []int, grumpy []int, X int) int {

	total := 0
	for k, v := range customers{
		if k < X{
			total += v
		}else{
			total += v * (1-grumpy[k])
		}

	}

	// add := 0
	res := total
	for i := X; i < len(grumpy); i++{
		add := customers[i] * grumpy[i] - customers[i-X] * grumpy[i-X]
		total += add
		if total > res{
			res = total
		}
	}

	return res

}
