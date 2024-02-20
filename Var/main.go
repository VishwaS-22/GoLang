package main

import "fmt"

const position string = "Student"

var clgName = "The Gopher's School"

func main() {
	var userName string = "Vishwa"
	fmt.Printf("The user's name is %v\n", userName)
	fmt.Printf("The type of user's name is %T\n", userName)

	var userDOB uint16 = 2003
	fmt.Printf("The user %v's birth year is %v\n", userName, userDOB)
	fmt.Printf("The type of user's DOB is %T\n", userDOB)

	userCGPA := 8.21
	fmt.Printf("The user %v's CGPA is %v\n", userName, userCGPA)
	fmt.Printf("The type of user's CGPA is %T\n", userCGPA)

	fmt.Printf("The desigination of user %v is %v\n", userName, position)
	fmt.Printf("The type of user's desigination is %T\n", position)

	fmt.Printf("The school of user %v is %v\n", userName, clgName)
	fmt.Printf("The type of user's school is %T\n", clgName)

}
