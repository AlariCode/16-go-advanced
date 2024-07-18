package main

import (
	"fmt"
	"time"
)

func main() {
	go printHi()
	go fmt.Println("Привет из main 2")
	go fmt.Println("Привет из main")
	time.Sleep(time.Second)
}

func printHi() {
	fmt.Println("Привет из gr")
}
