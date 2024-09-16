package main

import "fmt"

func main() {
    // Creating a map to store country names as keys and their capital cities as values
    capitals := make(map[string]string) //va

    // Adding elements to the map
    capitals["India"] = "New Delhi"
    capitals["USA"] = "Washington D.C."
    capitals["Japan"] = "Tokyo"

    // Accessing values from the map
    fmt.Println("Capital of India:", capitals["India"])
    fmt.Println("Capital of USA:", capitals["USA"])

    // Updating a value in the map
    capitals["USA"] = "New Washington"

    // Checking if a key exists in the map
    capital, exists := capitals["France"]
    if exists {
        fmt.Println("Capital of France:", capital)
    } else {
        fmt.Println("Capital of France: Not Found")
    }

    // Iterating over the map using range
    fmt.Println("\nAll countries and their capitals:")
    for country, capital := range capitals {
        fmt.Printf("%s: %s\n", country, capital)
    }

    // Deleting a key-value pair
    delete(capitals, "Japan")
    fmt.Println("\nAfter deleting Japan from map:", capitals)
}// OUT PUT--
// Capital of India: New Delhi
// Capital of USA: Washington D.C.
// Capital of France: Not Found

// All countries and their capitals:
// Japan: Tokyo
// India: New Delhi
// USA: New Washington

// After deleting Japan from map: map[India:New Delhi USA:New Washington]