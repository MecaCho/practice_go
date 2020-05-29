package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func ServerHandle1(w http.ResponseWriter, r *http.Request) {

	UsePlugin1()

	w.Write([]byte(fmt.Sprintf("hello, goroutine num %+v, time now: %+v\n", runtime.NumGoroutine(), time.Now())))
}

func ServerHandle2(w http.ResponseWriter, r *http.Request) {

	UsePlugin2()

	w.Write([]byte(fmt.Sprintf("hello, goroutine num %+v, time now: %+v\n", runtime.NumGoroutine(), time.Now())))
}

func ServerHandle3(w http.ResponseWriter, r *http.Request) {

	UsePlugin3()

	w.Write([]byte(fmt.Sprintf("hello, goroutine num %+v, time now: %+v\n", runtime.NumGoroutine(), time.Now())))
}

func StartServer() {
	// go func() {
	// 	for {
	// 		time.Sleep(1 * time.Second)
	//
	// 		fmt.Println("goroutine num: ", runtime.NumGoroutine(), time.Now())
	// 	}
	// }()
	// fmt.Println("goroutine num: ", runtime.NumGoroutine())

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/mss/service", ServerHandle1)
	mux.HandleFunc("/api/v1/mss/label", ServerHandle2)
	mux.HandleFunc("/api/v1/mss/port", ServerHandle3)
	fmt.Println("Start server....")
	http.ListenAndServe("", mux)
}
