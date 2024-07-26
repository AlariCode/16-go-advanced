package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	code := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getHttpCode(code)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(code)
	}()
	for res := range code {
		fmt.Printf("Код: %d\n", res)
	}
}

func getHttpCode(codeCh chan int) {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("Ошибка %s", err.Error())
	}
	codeCh <- resp.StatusCode
}
