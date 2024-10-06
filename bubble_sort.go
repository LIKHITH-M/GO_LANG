package main

import "fmt"

// BubbleSort function that sorts a slice of integers
func BubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                Swap(arr, j)  // Call Swap function to swap elements
            }
        }
    }
}

// Swap function that swaps two adjacent elements in the slice
func Swap(arr []int, i int) {
    arr[i], arr[i+1] = arr[i+1], arr[i]  // Swap the values
}

func main() {
    // First test: all positive numbers
    arr1 := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println("Unsorted array (all positive):", arr1)
    BubbleSort(arr1)
    fmt.Println("Sorted array (all positive):", arr1)

    // Second test: including negative numbers
    arr2 := []int{64, -34, 25, -12, 22, 11, -90}
    fmt.Println("Unsorted array (with negative numbers):", arr2)
    BubbleSort(arr2)
    fmt.Println("Sorted array (with negative numbers):", arr2)
}
