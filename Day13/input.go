package main

import "fmt"

func main(){

	var arr1 [5]int

	fmt.Println(arr1)

	fmt.Println( "Enter value for index 0.")
	fmt.Scanf( "%d", &arr1[0])

	fmt.Println( "Enter value for index 1.")
	fmt.Scanf( "%d", &arr1[1])

	fmt.Println( "Enter value for index 2.")
	fmt.Scanf( "%d", &arr1[2])

	fmt.Println( "Enter value for index 3.")
	fmt.Scanf( "%d", &arr1[3])

	fmt.Println( "Enter value for index 4.")
	fmt.Scanf( "%d", &arr1[4])


	fmt.Println(arr1)

	for i := 0 ; i < 5; i++ {
		fmt.Printf( "Enter value for index %d.\n", i)
		fmt.Scanf( "%d", &arr1[i])
	}

	fmt.Println(arr1)
}
