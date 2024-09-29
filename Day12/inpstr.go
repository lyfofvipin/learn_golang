package main

import "fmt"

func main() {

	var name string

	fmt.Printf("Enter Your Name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello %s!\n", name)
}
