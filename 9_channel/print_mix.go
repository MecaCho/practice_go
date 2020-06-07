package __channel

import (
	"fmt"
	"sync"
)

func MixPrint() {

	chan_n := make(chan bool)
	chan_c := make(chan bool, 1)
	done := make(chan struct{})

	chan_n <- true

	go func() {
		for i := 1; i < 11; i += 2 {
			<-chan_c
			fmt.Print(i)
			fmt.Print(i + 1)
			chan_n <- true
		}
	}()

	go func() {
		char_seq := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
		for i := 0; i < 10; i += 2 {
			<-chan_n
			fmt.Print(char_seq[i])
			fmt.Print(char_seq[i+1])
			chan_c <- true
		}
		done <- struct{}{}
	}()

	chan_c <- true
	<-done

}

func QuickSort(arr []int64) (res []int64) {

	var left []int64
	var right []int64

	if len(arr) < 2 {
		return arr
	}

	loc := arr[0]

	// fmt.Println(arr, loc)

	for k := range arr {
		if k == 0 {
			continue
		}
		// fmt.Println(arr[k])
		if arr[k] < loc {
			left = append(left, arr[k])
		} else {
			right = append(right, arr[k])
		}
	}

	// fmt.Println(left, right)

	l := QuickSort(left)
	for i, _ := range l {
		res = append(res, l[i])
	}
	res = append(res, arr[0])

	r := QuickSort(right)
	for i, _ := range r {
		res = append(res, r[i])
	}
	// fmt.Println(l, r)

	return res
}

func AddItem(arr *[]int64, item int64) {

	*arr = append(*arr, item)

}

func ChangeArrItem(arr [3]int64) {

	arr[0] = 100

}

func ChangeSliceItem(arr []int64) {

	arr[0] = 100

}

func PrintStr(a_chan, next_chan, done chan int, str string) {
	for i := 0; i < 100; i++ {
		select {
		case <-a_chan:
			fmt.Print(fmt.Sprintf("%s-%d", str, i))
			// fmt.Sprintf("%s-%d", str, i)
			next_chan <- i
		}
	}
	if str == "C" {
		done <- 1
	}
}

func PrintABC() {
	a_chan := make(chan int, 1)
	b_chan := make(chan int)
	c_chan := make(chan int)
	done := make(chan int, 1)
	a_chan <- 1
	go PrintStr(a_chan, b_chan, done, "A")
	go PrintStr(b_chan, c_chan, done, "B")
	go PrintStr(c_chan, a_chan, done, "C")

	<-done
	// fmt.Println("\nend", <-done)
}

func PrintStrWaitGroup(a_chan, next_chan chan int, str string, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		select {
		case v, ok := <-a_chan:
			fmt.Println(v, ok)
			// fmt.Print(fmt.Sprintf("%s-%d", str, i))
			fmt.Sprintf("%s-%d", str, i)
			next_chan <- i
		}
	}
	wg.Done()
}

func PrintWaitGroup() {
	a_chan := make(chan int, 1)
	b_chan := make(chan int, 1)
	c_chan := make(chan int, 1)
	wg := sync.WaitGroup{}

	wg.Add(3)

	go PrintStrWaitGroup(a_chan, b_chan, "A", &wg)
	go PrintStrWaitGroup(b_chan, c_chan, "B", &wg)
	go PrintStrWaitGroup(c_chan, a_chan, "C", &wg)
	a_chan <- 1
	wg.Wait()
	// fmt.Println("\nend")
}
