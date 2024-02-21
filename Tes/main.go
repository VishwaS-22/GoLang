package main

import "fmt"

// Input a year and find whether it is a leap year or not.
/*
func main() {

	year := 2023
	fmt.Println("Year: ", year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Printf("The year %v is Leap year", year)
	} else {
		fmt.Printf("The year %v is Not a Leap year", year)
	}

}
*/

// Take two numbers and print the sum of both.
/*
func main() {
	v := 11
	s := 11
	res := v + s
	fmt.Printf("Sum of v and s is: %v", res)
}*/

//Take a number as input and print the multiplication table for it.

/*func main() {
	v := 10
	fmt.Printf("Multiplication table for the input value %v is\n", v)
	mTable(v)
}
func mTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d is %d\n", i, n, i*n)
	}
}*/
// Take 2 numbers as inputs and find their HCF and LCM.

func main() {
	fmt.Println("Printing HCM and LCM")
	n := 10
	m := 5

	if n < m {
		v := n
	} else {
		v := m
	}
	if n > m {
		s := n
	} else {
		s := m
	}

	for i := 0; i <= v; i++ {
		if i%v == 0 && i%s == 0 {
			fmt.Print("LCM", i)
		}
	}

}
