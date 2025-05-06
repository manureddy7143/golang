package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog"

	// Basic string operations

	// TODO: Length of string
	fmt.Println(len(s))

	// TODO: iterate over each character
	for _, ch := range s {
		fmt.Print(string(ch), ",")
	}
	fmt.Println()

	// TODO: Using operators < > == !=

	fmt.Println("dog" < "cat")
	fmt.Println("dog" < "horse")
	fmt.Println("dog" == "Dog")

	// TODO: Comparing two strings

	result := strings.Compare("dog", "cat")
	fmt.Println(result)
	result = strings.Compare("dog", "dog")
	fmt.Println(result)

	// TODO: EqualFold tests using Unicode case-folding
	fmt.Println(strings.EqualFold("dog", "Dog"))

	// TODO: ToUpper, ToLower, Title
	s1 := strings.ToUpper(s)
	s2 := strings.ToLower(s)
	s3 := strings.Title(s)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
