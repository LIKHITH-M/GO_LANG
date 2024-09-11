package main 

import "fmt"

func main(){

	var name string

	fmt.Println("enter the name:")
//function to read the input
//takes only continous char till it encouters the end line character or space even if you write enter using double quotes
	fmt.Scanln(&name)

	fmt.Printf("hello, %s!",name)

}