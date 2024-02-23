package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Slices here")
	/*
		var arr1 []int
		arr1 = append(arr1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		fmt.Println(arr1)
		fmt.Println(len(arr1))

		i := 4
		arr1 = append(arr1[:i], arr1[i+1:]...)
		fmt.Println(arr1)
	*/

	read := bufio.NewReader(os.Stdin) // read input
	in, _ := read.ReadString('\n')    // read till new line arrives

	arr := strings.Fields(in) // split string as individual tokens

	var arr1 []int //slice

	for _, val := range arr {

		n, err := strconv.Atoi(val)

		if err != nil {
			fmt.Println(err)
			return
		}
		arr1 = append(arr1, n)

	}
	for id, val1 := range arr1 {
		fmt.Printf("%v:%v\n", id, val1) // printing elemtent with Index.
	}
}
