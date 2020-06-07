package weekly_192

import (
	"sort"
	"math"
)

func getStrongest(arr []int, k int) []int {

	sort.Ints(arr)
	mid := arr[(len(arr)- 1)/2]
	sort.Slice(arr, func(i, j int) bool {
		if math.Abs(float64(arr[i] - mid)) > math.Abs(float64(arr[j] - mid)) || (math.Abs(float64(arr[i] - mid)) == math.Abs(float64(arr[j] - mid)) && arr[i] > arr[j]){
			return true
		}
		// if  math.Abs(float64(arr[i] - mid)) == math.Abs(float64(arr[j] - mid)){
		//     return arr[i] > arr[j]
		// }
		return false

	})
	// fmt.Println(arr)
	return arr[:k]

}