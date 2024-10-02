package main

import "fmt"

func main(){

	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	b := a[3:7]

	fmt.Println(a)
	fmt.Println(b)

}
