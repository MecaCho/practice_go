package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "get url (%s) error: %s\n", url, err.Error())
			os.Exit(1)
		}
		if resp.Body != nil {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "get url (%s) reading error: %s\n", url, err.Error())
				os.Exit(1)
			}
			fmt.Printf("%s", body)
		}
	}
}
