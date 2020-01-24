package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

var logg *log.Logger

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()

}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	sil := []int{1, 2, 3}
	a := 1
	a++
	a--
	if true {
		fmt.Println(sil)
	}

	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func main() {

	go func() {
		log.Println(http.ListenAndServe("127.0.0.1:6060", nil))
	}()

	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("down")
}
