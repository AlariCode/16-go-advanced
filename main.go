package main

import "fmt"

func sumPart(arr []int, ch chan int) {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	ch <- sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numGoroutines := 3
	ch := make(chan int, numGoroutines)

	partSize := len(arr) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := start + partSize
		go sumPart(arr[start:end], ch)
	}

	totalSum := 0
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-ch
	}
	fmt.Println("Total sum: ", totalSum)
}
