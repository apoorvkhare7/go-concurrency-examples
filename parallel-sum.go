package main

import (
	"fmt"
	"sync"
)

func SUM1() int {
	n := 1000000
	messages := make(chan int)
	var wg sync.WaitGroup

	wg.Add(n)

	for i := 1; i <= n; i++ {
		go func(i int) {
			defer wg.Done()
			messages <- i

		}(i)
	}

	var sum int
	go func() {
		for i := range messages {
			sum = sum + i
		}
	}()

	//go func() {
	//	wg.Wait()
	//	close(messages)
	//}()
	wg.Wait()

	return sum
}

func main() {
	s := SUM1()

	fmt.Println(s)
}
