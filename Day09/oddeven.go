package main

import ("fmt")

func main() {

	value := 0

	switch value > 0 {

	case true:
		fmt.Printf("%d is a +ve value.\n", value)
	case false:
		fmt.Printf("%d is a -ve value.\n", value)
	}

}
