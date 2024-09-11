package main

import "fmt"
//here iota assigns the different values to diff constants, assigning to first value is enough.
const (
    Zero = iota // 0
    One         // 1
    Two         // 2
    Three       // 3
)
const (
	_ = iota // ignore first value by assigning to blank identifier
	Bit0 = 1 << iota
	Bit1
	Bit2
	Bit3
)

func main() {
    fmt.Println("Basic Enumeration:")
    fmt.Printf("Zero: %d\n", Zero)
    fmt.Printf("One: %d\n", One)
    fmt.Printf("Two: %d\n", Two)
    fmt.Printf("Three: %d\n", Three)
fmt.Println("\n----------------------\n")
// iota used for bit shifting
	fmt.Println("Bit Masks:")
		fmt.Printf("Bit0: %b\n", Bit0) // 0001
		fmt.Printf("Bit1: %b\n", Bit1) // 0010
		fmt.Printf("Bit2: %b\n", Bit2) // 0100
		fmt.Printf("Bit3: %b\n", Bit3) // 1000
	
}
