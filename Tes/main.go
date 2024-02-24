package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Slices here")
	/*
		var arr1 []int
		arr1 = append(arr1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		fmt.Println(arr1)
		fmt.Println(len(arr1))

		i := 4
		arr1 = append(arr1[:i], arr1[i+1:]...)
		fmt.Println(arr1)
	*/

	/*read := bufio.NewReader(os.Stdin) // read input
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
	}*/

	var n int
	fmt.Println("Enter a limit for loop:")
	fmt.Scan(&n)
	var n1 int
	//var arr []int
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&n1)
		//arr = append(arr, n1)
		arr[i] = n1 // adding elements to the slice
	}
	fmt.Println(arr)

	//Finding max
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Println("The max val of slice is: ", max)

	//Finding min of slice
	min := arr[0]
	for _, i := range arr[1:] {
		if i < min {
			min = i
		}
	}
	fmt.Println("The min val of slice is: ", min)
}
