package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Gonna handle time in Golang!")

	pTime := time.Now()
	fmt.Println(pTime)

	fmt.Println(pTime.Format("01-02-2006 15:04:05 Monday"))

	cDate := time.Date(2023, time.September, 10, 10, 55, 0, 0, time.UTC)
	fmt.Println(cDate)
	fmt.Println(cDate.Format("01-02-2006 Monday"))
}
