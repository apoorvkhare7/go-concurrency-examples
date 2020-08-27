package main

import "fmt"

// function to add an array of numbers.
func sum(a int, b int, c chan int) {
	sum := 0
	for i:=a; i<=b;i++ {
		sum += i
	}
	// writes the sum to the go routines.
	c <- sum // send sum to c
}

func main() {
	n := 100000000
	const chan_count = 100
	var chans [chan_count]chan int
	for i := range chans {
		chans[i] = make(chan int)
	}
	for i:=0; i<chan_count;i++{
		go sum(i*(n/chan_count)+1, (n/chan_count)*(i+1), chans[i])
		// spin up a goroutine.
	}
	var ans [chan_count]int
	result := 0
	for i:=0; i<chan_count;i++{
		ans[i] = <-chans[i]
		fmt.Println(ans[i])
		result += ans[i]
	}

	fmt.Println(result)
}