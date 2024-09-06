package main

import ("fmt")

func main(){

	var i int
	var j int

	fmt.Printf("Input via Scan: ")
	fmt.Scanf("%d-%d", &i, &j)
	fmt.Printf("Value of %d is %d.\n", i, j)

}