package main 

import "fmt"

func main(){
	
	arr := [2][3]int{
		{1,2,3},
		{4,5,6},
	} //comma's are required after each rows.

	fmt.Println(arr)
}// output [[1 2 3][4 5 6]]