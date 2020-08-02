package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get url (%s) error: %s\n", url, err.Error())
		os.Exit(1)
	}
	if resp.Body == nil {
		fmt.Fprintf(os.Stderr, "get url (%s) error: %s\n", url, "response nil")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get url (%s) reading error: %s\n", url, err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s", body)
	ch <- fmt.Sprintf("%.6f seconds %7d %s", time.Since(start).Seconds(), len(body), url)

}

func main() {

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.6f secondes elapsed\n", time.Since(start).Seconds())
}
