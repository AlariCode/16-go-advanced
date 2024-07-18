package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// 10 конкурентных запросов на GET по адресу google.com
// Вывести в консоль 10 StatusCode

func main() {
	t := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getHttpCode()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Since(t))
}

func getHttpCode() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("Ошибка %s", err.Error())
	}
	fmt.Printf("Код: %d\n", resp.StatusCode)
}
