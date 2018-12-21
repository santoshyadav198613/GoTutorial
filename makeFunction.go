package main

import "fmt"

func makeDemo() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	var s []int
	printSlices(s)

	s = append(s, 0)
	printSlices(s)

	s = append(s, 1)
	printSlices(s)

	s = append(s, 2, 3, 4)
	printSlices(s)

}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
