package main

import (
    "fmt"
)

func main() {
    // Elemtents
    ets := []int{23, 54, 36, 20, 14, 100, 1}

    // Initializeing max and min with the elemtent one of the array
    m := ets[0]
    n := ets[0]

    // Iterate through the elements
    for _, element := range ets {
        // Update max if the current element is greater
        if element > m {
            m = element
        }
        // Update min if the current element is smaller
        if element < n {
            n = element
        }
    }

    // Print the maximum and minimum values
    fmt.Println("Maximum:", m)
    fmt.Println("Minimum:", n)
}
