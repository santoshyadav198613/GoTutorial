package main

import "fmt"

var powVal = []int{1, 2, 4, 8, 16, 32, 64, 128}

func rangeDemo() {
	for i, v := range powVal {
		fmt.Printf("2**%d=%d\n", i, v)
	}

	power := make([]int, 10)
	for i := range power {
		power[i] = 1 << uint(i)
	}

	for _, value := range power {
		fmt.Printf("%d\n", value)
	}
}
