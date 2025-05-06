package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10.0

	// TODO: Absolute value

	fmt.Println(math.Abs(x))
	fmt.Println(math.Abs(-x))

	// TODO: Pow (x^y) and Exp (e^x)
	fmt.Println(math.Pow(x, 2.0))
	fmt.Println(math.E, math.Exp(x))

	// TODO: Trigonometry and other functions

	fmt.Println(math.Cos(math.Pi))

	fmt.Println(math.Sin(math.Pi) / 2)
	fmt.Println(math.Tan(math.Pi) / 2)

	// TODO: Square and cube roots
	fmt.Println(math.Sqrt(25))
	fmt.Println(math.Cbrt(125))

	// TODO: calculate the Hypotenuse of a right triangle

}
