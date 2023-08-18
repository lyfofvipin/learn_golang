package main

import ("fmt")

func main(){
	fmt.Printf("This is an int value %d.\n", 7)
	fmt.Printf("This is an float value %g.\n", 45.6)
	fmt.Printf("This is an string value without quote %s.\n", "Vipin")
	fmt.Printf("This is an string value with quote %q.\n", "vipin")
	fmt.Printf("This is an boolean value %t.\n", false)
	fmt.Printf("Type of first value %T, second value %T.\n", 7, "Vipin")
	fmt.Printf("Print int  %v, print str %v print float.\n", 7, "Vipin", 45.6)
}