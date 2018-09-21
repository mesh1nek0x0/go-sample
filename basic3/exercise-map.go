package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	tmp := strings.Fields(s)
	m := make(map[string]int)
	for i := 0; i < len(tmp); i++ {
		// Fieldsの結果の中にいくつあるのか？を調べる必要がある
		m[tmp[i]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
