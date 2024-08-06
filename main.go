package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ping(url string, respCh chan int, errCh chan error) {
	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	respCh <- resp.StatusCode
}

func main() {
	path := flag.String("file", "url.txt", "path to URL file")
	flag.Parse()
	file, err := os.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}
	urlSlice := strings.Split(string(file), "\n")
	respCh := make(chan int)
	errCh := make(chan error)
	for _, url := range urlSlice {
		go ping(url, respCh, errCh)
	}
	for range urlSlice {
		select {
		case err := <-errCh:
			fmt.Println(err)
		case res := <-respCh:
			fmt.Println(res)
		}
	}
}
