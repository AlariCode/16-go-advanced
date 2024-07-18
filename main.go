package main

import (
	"fmt"
	"net/http"
	"time"
)

// 10 конкурентных запросов на GET по адресу google.com
// Вывести в консоль 10 StatusCode

func main() {
	t := time.Now()
	for i := 0; i < 10; i++ {
		go getHttpCode()
	}
	time.Sleep(time.Millisecond * 1100)
	fmt.Println(time.Since(t))
}

func getHttpCode() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("Ошибка %s", err.Error())
	}
	fmt.Printf("Код: %d\n", resp.StatusCode)
}
