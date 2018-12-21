package main

import "fmt"

func pointer() {
	i, j := 42, 2700

	p := &i         //pointer to i
	fmt.Println(*p) //reading i through pointer
	*p = 21         //change value of i through pointer
	fmt.Println(i)  //get new value of i

	p = &j         //pointer to j
	*p = *p / 37   // divide j with pointer
	fmt.Println(j) //new value to j

}
