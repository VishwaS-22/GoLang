package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a Value:")

	userInput, _ := read.ReadString('\n')
	fmt.Print("This is the value, ", userInput)
	fmt.Printf("The type of value is %T", userInput)
}
