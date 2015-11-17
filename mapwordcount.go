package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	split := strings.Fields(s)
	for i := 0; i < len(split); i++ {
		word := split[i]
		m[word]++;
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
