package main

import "fmt"

func main() {
	naturals := make(chan int)
	sums := make(chan int)
	// Counter
	go func() {
		for x := 0; x <= 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	// Squarer
	result := 0
	go func() {
		for x := range naturals {
			result += x
			sums <- result
		}
		close(sums)
	}()
	// Printer (in main goroutine)
	for x := range sums {
		fmt.Println(x)
	}

}
