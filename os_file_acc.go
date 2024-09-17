package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// File name to open
	fileName := "names.txt"

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file and print first name/last name pairs
	for scanner.Scan() {
		line := scanner.Text()

		// Split each line into first and last name
		names := strings.Split(line, " ")
		if len(names) >= 2 {
			firstName := names[0]
			lastName := names[1]

			// Print first and last name
			fmt.Printf("First Name: %s, Last Name: %s\n", firstName, lastName)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
// First Name: John, Last Name: Doe
// First Name: Alice, Last Name: Smith
// First Name: Bob, Last Name: Johnson
