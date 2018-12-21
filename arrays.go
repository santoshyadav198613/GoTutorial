package main

import "fmt"

func arrayDemo() {

	names := [4]string{"Amit", "Sumit", "Ram", "Shyam"}
	fmt.Println(names)

	i := names[0:2]
	j := names[1:3]

	fmt.Println(i, j)
	j[0] = "Test"
	fmt.Println(i, j)

	fmt.Println(names)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	var s []int = primes[1:4] //slice
	fmt.Println(s)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(t)
	fmt.Println(len(t), cap(t)) // length and capacity

	t = t[:2]

	fmt.Println(t)
	fmt.Println(len(t), cap(t))

	var empty []int
	fmt.Println(empty, len(empty), cap(empty))

	if empty == nil {
		fmt.Println("nil!")
	}
}
