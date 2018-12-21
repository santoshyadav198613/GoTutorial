package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	x := strings.Fields(s)	
	for i := 0; i < len(x); i++ {
		m[x[i]] += 1
		fmt.Println(m[x[i]])
	}

	return m
}

func wordCount() {
	wc.Test(WordCount)
	//WordCount("Hello World")
}
