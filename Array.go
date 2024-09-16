package main 

import "fmt"

func main(){
	fmt.Println("array program")
	
	//basic type of initialising and declaring an array.
	var arr = [3]int{1,2,3} //the value in the square braces represent the size.
	fmt.Println(arr) //printing complete array.
	//output [1,2,3]


	//using the colon operator.
	arr2 := [4]int{2,4,5} //here the 3rd index by default initialized to 0.
	fmt.Println(arr2) // output [2,4,5,0]

	//using the period operator for size.
	arr3 := [...]int{2,34,5,6,62,2,4} //size is assigned wrt to number of array elements. Here ... is a keyword
	fmt.Println(arr3)
}