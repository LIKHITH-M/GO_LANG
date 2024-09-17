package main
import (
	"encoding/json"
	"fmt"
)
func main() {
	// Variables to store input
	var name, address string
	// Reading input from the user
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your address: ")
	fmt.Scanln(&address)
	// Creating a map to hold the name and address
	data := map[string]string{
		"name":    name,
		"address": address,
	}
	// Marshalling the map into a JSON object
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Printing the JSON object
	fmt.Println(string(jsonData))
}  // output
// Enter your name: likhith
// Enter your address: channapatana
// {
//   "address": "channapatana",
//   "name": "likhith"
// }
