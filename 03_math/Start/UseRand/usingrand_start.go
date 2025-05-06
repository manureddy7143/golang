package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	// Seed the random number generator
	//rand.Seed(time.Now().UnixNano())

	// TODO: Shuffle a sequence of values
	const numstring = "one two three four five six"
	// words := strings.Fields(numstring)
	// rand.Shuffle(len(words), func(i, j int) {
	// 	words[i], words[j] = words[j], words[i]
	// })
	// fmt.Println(words)

	// TODO: Generate a random string
	const upletters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const lowletters = "abcdefghijklmnopqrstuvwxyz"
	const digits = "0123456789"
	const specialchars = "~=+%^*()[]{}!@#$?|"
	const allchars = upletters + lowletters + digits + specialchars

	var sb strings.Builder

	length := 10

	for i := 0; i < length; i++ {
		sb.WriteRune(rune(allchars[rand.Intn(len(allchars))]))
	}

	fmt.Println(sb.String())

	// TODO: Generate random string with rules

	buf:= make([byte],length)
	buf[0]= digits
}
