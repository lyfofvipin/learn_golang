package main

import ("fmt")

func main(){

	var i int
	var j int

	fmt.Printf("Input via Scan: ")
	fmt.Scanln(&i, &j)
	fmt.Printf("Value of %d is %d.\n", i, j)

}