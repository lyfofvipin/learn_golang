package main

import (
	"fmt"
)

func main (){

	var arr1 [5]int
	arr2 := [5]int{ 10, 20, 30, 40, 50 }

	var arr3 = [...]int{ 1, 2, 3, 4, 5 }
	arr4 := [...]int{ 10, 20, 30, 40, 50 }

	var arr5 = [...]string{ "A", "a", "b", "B", `c`, `C` }
	arr6 := [...]rune{ 'a', '4', '-', '+' }

	arr7 := [...]bool { true }
	arr8 := [...]bool { true }

	fmt.Println(arr1, arr2, arr3, arr4, arr5, arr6, arr7, arr8 )

}
