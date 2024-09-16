package main

import "fmt"

// Define the Person struct
type Person struct {
    firstName string
    lastName  string
    age       int
}

// Method associated with the Person struct
func (p Person) Greet() string {
    return "Hello, my name is " + p.firstName + " " + p.lastName
}

func main() {
    // Create an instance of Person
    person := Person{firstName: "Alice", lastName: "Smith", age: 30}

    // Call the Greet method
    fmt.Println(person.Greet())
}
// output
// Hello, my name is Alice Smith