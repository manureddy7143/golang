package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	S, _ := reader.ReadString('\n')
	fmt.Println(S)
}
