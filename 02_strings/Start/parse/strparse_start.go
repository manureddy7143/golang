package main

import (
	"fmt"
	"strconv"
	//"strconv"
)

func main() {
	sampleint := 100
	samplestr := "250"

	// This does a character conversion, not a numerical one
	// newstr := string(sampleint)
	// fmt.Println("Result of using string():", newstr)

	// The strconv package contains a variety of functions for parsing and formatting
	// numbers, values, and strings

	// Convert an integer to string
	s := strconv.Itoa(sampleint)
	fmt.Printf("%T ,%v \n", s, s)

	// Convert a string to integer
	i, err := strconv.Atoi(samplestr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T ,%v \n", i, i)

	// Other parse functions
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)

	f, _ := strconv.ParseFloat("3.14158", 64)
	fmt.Println(f)

	// Other format functions
	s = strconv.FormatBool(true)
	fmt.Printf("%T ,%v", s, s)

}
