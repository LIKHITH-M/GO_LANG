package main

import "fmt"

func main() {
    // Define a slice of integers
    numbers := []int{10, 20, 30, 40, 50} //the difference between slice and array is we mention the size for array but not for slice!!!!!.

    // Print the length and capacity of the slice
    fmt.Printf("Length: %d, Capacity: %d\n\n", len(numbers), cap(numbers))

    // Use range to iterate over the slice
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
