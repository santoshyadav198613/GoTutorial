package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8,dy)

	for x:=0;x< dy;x++ {
		image[x] = make([]uint8,dx)
		for y:=0;y< dx; y++ {
			image[x][y]= uint8((x^y)* (x^y)) 
		}
	}
	return image
}

func sliceExercise() {
	pic.Show(Pic)
}