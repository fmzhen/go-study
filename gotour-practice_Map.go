package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	res := make(map[string]int)
	for _, word := range words {
		_, ok := res[word]
		if ok {
			res[word]++
		} else {
			res[word] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
