package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"syscall"
	"time"
)

func IsNotExist(err error) bool {
	if pe, ok := err.(*os.PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responsing(%s); retrying ...", err)
		time.Sleep(time.Second << uint(tries)) // 指数退避策略
	}
	return fmt.Errorf("swerver %s failed to response afater %s", url, timeout)

}

func helloworld(w http.ResponseWriter, r *http.Request) {
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
		log.Fatalf("Site is down : %v\n", err)
	}

	// w.Write([]byte("hello world\n"))
	io.WriteString(w, "hello world\n")
	// _, err := os.Open("xxx")
	// fmt.Println(os.IsNotExist(err))
}

func main() {
	runtime.GOMAXPROCS(1)
	http.HandleFunc("/", helloworld)
	http.ListenAndServe(":3000", nil)
}
