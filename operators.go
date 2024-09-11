package main

import (
    "fmt"
	"strings"  //importing string package for string operations.
   
)

func main(){
		// Arithmetic operations.
		a:=50
		b:=25

		fmt.Println("addition of a and b: %d",a+b)
		fmt.Println("substraction of a a by b : %d",a-b)
		fmt.Println("multiplication of a and b: %d",a*b)
		fmt.Println("division of a by b : %d",a/b)
		fmt.Println("modulus of a by b : %d",a%b)
		fmt.Println("\n--------------------------------------\n")

		// String operations.

		str := "Hello,bot!"
		fmt.Println("orginal string: %s",str)
		fmt.Println("upper case of org: %s",strings.ToUpper(str))
		fmt.Println("lower case of org : %s",strings.ToLower(str))
		fmt.Println("replaceing bot with pro : %s",strings.ReplaceAll(str,"bot","pro"))
}