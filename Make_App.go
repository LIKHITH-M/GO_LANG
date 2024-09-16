package main

import "fmt"

func main() {
    // Using make() to create an empty slice of integers with length 0 and capacity 3
    numbers := make([]int, 0, 3)

    // Display the initial slice, its length, and capacity
    fmt.Printf("Initial slice: %v, Length: %d, Capacity: %d\n", numbers, len(numbers), cap(numbers))

    // Append elements to the slice
    numbers = append(numbers, 10)
    numbers = append(numbers, 20)
    numbers = append(numbers, 30)

    // Display the slice after appending elements
    fmt.Printf("After appending: %v, Length: %d, Capacity: %d\n", numbers, len(numbers), cap(numbers))

    // Append another element, causing the capacity to increase
    numbers = append(numbers, 40)

    // Display the slice after appending more than its initial capacity
    fmt.Printf("After appending beyond capacity: %v, Length: %d, Capacity: %d\n", numbers, len(numbers), cap(numbers)) 
	// %v format specifier is a versatile placeholder used with the fmt package to print values of any type in their default format. 
}
// Initial slice: [], Length: 0, Capacity: 3
// After appending: [10 20 30], Length: 3, Capacity: 3
// After appending beyond capacity: [10 20 30 40], Length: 4, Capacity: 6
