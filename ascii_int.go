package main
import (
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi("10") // Atoi: Converts an ASCII-encoded string representation of a number to an integer
	y := i * 2
	fmt.Println(y) // output is 20
  }
  