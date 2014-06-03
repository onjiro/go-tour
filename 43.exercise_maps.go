package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	counts := make(map[string]int)

	for _, w := range words {
		counts[w] = counts[w] + 1
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
