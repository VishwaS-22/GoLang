package main

import "fmt"

func main()  {
	fmt.Println("Pointers in Golang!")

	var v *int
	fmt.Println("The value of v is ", v)

	myNum := 22
	var p = &myNum
    *p = *p+2 
	p1 := *p+1
	
	fmt.Println("Address of myNum is ", p)
	fmt.Println("Value of myNum is ", *p)
	fmt.Println("Value of ptr is ", p1)
}