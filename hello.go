package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
	"time"
)

// func add(num1 int,num2 int) int {
//     return num1+ num2;
// }

// var c, python, Java bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
var c, python, Java = true, false, "No!"

func add(num1, num2 int) int {
	return num1 + num2
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func typeInference() {
	v := 42.5
	fmt.Printf("v is type of %T\n", v)
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func forLoop() {
	// sum := 0
	sum := 1
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }

	for sum < 100 { // for loop as while loop
		sum += sum
	}
	fmt.Printf("Total is: %v\n", sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func switchStatement() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Print("LINUX.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switchCase() {
	fmt.Println("When's Thursday?")

	today := time.Now().Weekday()

	switch time.Thursday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning.")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon.")
	default:
		fmt.Println("Good Evening.")
	}
}

func main() {
	switchNoCondition()
	switchCase()
	switchStatement()
	fmt.Println(
		pow(2, 2, 10),
		pow(2, 3, 20),
		pow(3, 3, 10),
	)
	fmt.Println(sqrt(2), sqrt(-4))
	forLoop()
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	const Pi = 3.14
	fmt.Print("Happy", Pi, "Day")
	var i, j int = 1, 2
	k := 3 //short variable declaration
	typeConversion()
	typeInference()
	fmt.Println(i, j, k, c, python, Java)
	fmt.Println(split(10))
	a, b := swap("Go is", "Great")
	fmt.Println(a, b)
	defer fmt.Println(add(2, 4)) //defers the execution till other funcions are executed
	fmt.Printf("hello, world\n")
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
