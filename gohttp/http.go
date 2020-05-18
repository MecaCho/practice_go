package gohttp

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func ServerHandle(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(fmt.Sprintf("hello, goroutine num %+v, time now: %+v\n", runtime.NumGoroutine(), time.Now())))
}

func ServerHandle1(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(fmt.Sprintf("hello, goroutine num %+v, time now: %+v\n", runtime.NumGoroutine(), time.Now())))
}

func Server() {
	// go func() {
	// 	for {
	// 		time.Sleep(1 * time.Second)
	//
	// 		fmt.Println("goroutine num: ", runtime.NumGoroutine(), time.Now())
	// 	}
	// }()
	// fmt.Println("goroutine num: ", runtime.NumGoroutine())

	mux := http.NewServeMux()
	mux.HandleFunc("/", ServerHandle)
	mux.HandleFunc("/1", ServerHandle1)
	fmt.Println("Start server....")
	http.ListenAndServe("", mux)
	fmt.Println("goroutine num: ", runtime.NumGoroutine())

}
