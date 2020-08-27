package main

import (
	"fmt"
	"testing"
)

func SUM(a int, b int, ch chan int) {
	total := 0
	for i := a; i <= b; i++ {
		total += i
	}
	ch <- total
}

func chSumBuffer() (result_sum int) {
	n := 100000000
	const chan_count = 100

	ch := make(chan int, chan_count)
	for i := 0; i < chan_count; i++ {
		go SUM(i*(n/chan_count)+1, (n/chan_count)*(i+1), ch)
	}
	result_sum = 0
	for i := 0; i < chan_count; i++ {
		x := <-ch
		result_sum += x
	}

	return
}

func BenchmarkchSum(b *testing.B) {
	result := chSumBuffer()
	fmt.Println(result)
}

func main() {
	fmt.Println(testing.Benchmark(BenchmarkchSum))

}
