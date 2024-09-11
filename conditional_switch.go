    package main
	import "fmt"

	func main(){
	// 1. Conditional Statements
    fmt.Println("\nConditional Statements Example:")
    
    age := 18
    if age < 18 {
        fmt.Println("Underage")
    } else if age == 18 {
        fmt.Println("Exactly 18")
    } else {
        fmt.Println("Overage")
    }

    // 2. Switch Case Example
    fmt.Println("\nSwitch Case Example:")

    day := 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Another day")
    }
}