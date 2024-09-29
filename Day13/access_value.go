package main

import "fmt"

func main(){

	arr1 := [5]string{ "My", "Name", "Is", "Vipin", "Yadav" }

	fmt.Println(arr1[3])

	// To get the specific value from the string we can use it's index
	fmt.Printf( "%c\n", arr1[0][0] )

}