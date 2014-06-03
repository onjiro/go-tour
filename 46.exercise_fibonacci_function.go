package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	memo := []int{}
	return func() int {
		var value int
		if len(memo) <= 1 {
			value = 1
		} else {
			value = memo[len(memo)-2] + memo[len(memo)-1]
		}
		memo = append(memo, value)
		return value
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
