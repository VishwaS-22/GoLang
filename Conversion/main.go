package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Read input from the user.")

	read := bufio.NewReader(os.Stdin)

	input, _ := read.ReadString('\n')
	fmt.Println("This is the input", input)

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add one to input:", num+1)
	}
}
