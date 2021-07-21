package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, s := range strings.Fields(s) {
		m[s]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
