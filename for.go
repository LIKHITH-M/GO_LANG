package main

import "fmt"

func main() {
    // 1. For Loop Examples
    fmt.Println("For Loop Examples:")

    // Simple for loop
    for i := 0; i < 3; i++ {
        fmt.Printf("Simple for loop, iteration: %d\n", i)
    }

    // For loop with a slice
    numbers := []int{10, 20, 30}
    for _, num := range numbers {
        fmt.Printf("Iterating over a slice: %d\n", num)
    }

    // For loop with index and value
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}