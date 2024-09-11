package main

import( "fmt"
		"strconv"
)

func main() {
    // Initial values
    var intVal int = 42
    var floatVal float64 = 3.14

    // Type casting
    var floatFromInt float64 = float64(intVal)
    var intFromFloat int = int(floatVal)

    fmt.Println("Type Casting:")
    fmt.Printf("Int value: %d\n", intVal)
    fmt.Printf("Float value: %.2f\n", floatVal)
    fmt.Printf("Float from int: %.2f\n", floatFromInt)
    fmt.Printf("Int from float: %d\n", intFromFloat)

	fmt.Println("\n---------------------\n")

	    // Integer to string conversion
		num := 6789
		str := strconv.Itoa(num)
	
		fmt.Println("Integer to String Conversion:")
		fmt.Printf("Integer: %d\n", num)
		fmt.Printf("String: %s\n", str)
}
