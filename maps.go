package main

import "fmt"

type VertexVal struct {
	Lat, Long float64
}

// var m map[string]VertexVal

var m = map[string]VertexVal{
	"Hello": { //if type is defined it can be removed while assigning values
		40.05, -74.87,
	},
	"World": VertexVal{
		35.667, -155.5464,
	},
}

func mapDemo() {
	// m = make(map[string]VertexVal)
	// m["Hello"] = VertexVal{40.0, -74.87}
	// fmt.Println(m["Hello"])
	fmt.Println(m)
	mutatingDemo()
}

func mutatingDemo() {
	m := make(map[string]int)

	m["Answer"] = 45
	fmt.Printf("The value is %d\n", m["Answer"])

	m["Answer"] = 48 //updated value
	fmt.Printf("The value is %d\n", m["Answer"])

	delete(m, "Answer")
	fmt.Printf("The value is %d\n", m["Answer"])

	v, ok := m["Answer"]
	fmt.Printf("The value is %d Present %v\n", v, ok)

}
