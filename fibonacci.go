package main

import (
	"fmt"
	"time"
)

func fibBottomUp(n int) int {
	memo := make([]int, n + 1)
	memo[0] = 1
	memo[1] = 1


	if n == 1 || n == 2 {
		return 1
	}

	for i := 2; i < n; i++ {
		memo[i] = memo[i - 1] + memo[i - 2]
	}

	return memo[n - 1]
}

func fibMemoized(n int, memo []int) int {
	var result int
	index := n - 1
	if (memo[index] != 0) {
		result = memo[index]
	} else {
		result = fibMemoized(n - 1, memo) + fibMemoized(n - 2, memo)
		memo[index] = result
	}

	return result
}

func fib(n int) int {
	var result int
	
	if n == 1 || n == 2 {
		result = 1
	} else {
		result = fib(n - 1) + fib(n - 2)
	}

	return result
}

func elapsed(start, end time.Time) float64 {
	return end.Sub(start).Seconds() * 1000
}

func showResult(algorithm string, n, r int, elapsed float64) {
	fmt.Printf("%s(%d): %d, time: %f milliseconds\n", algorithm, n, r, elapsed)
}

func main(){
	n := 1000
	var r int
	var t1, t2 time.Time

	// t1 := time.Now()
	// r := fib(n)
	// t2 := time.Now()
	
	// showResult("fib", n, r, elapsed(t1, t2))

	t1 = time.Now()
	memo := make([]int, n)
	memo[0] = 1
	memo[1] = 1
	r = fibMemoized(n, memo)
	t2 = time.Now()
	
	showResult("fibMemoized", n, r, elapsed(t1, t2))

	t1 = time.Now()
	r = fibBottomUp(n)
	t2 = time.Now()
	
	showResult("fibBottomUp", n, r, elapsed(t1, t2))
}