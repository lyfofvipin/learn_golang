package main

import "fmt"

func main(){
	
	var arr1 [2][2]int
	var arr2 = [2][2]int{
		{ 1, 2 },
		{ 3, 4 },
	}

	var arr3 = [3][3]int{
		{ 1, 2, 3},
		{ 1, 2, 3},
		{ 1, 2, 3},
	}

	var arr4 = [4][5]int{
		{ 1, 2, 3, 4, 5},
		{ 1, 2, 3, 4, 5},
		{ 1, 2, 3, 4, 5},
		{ 1, 2, 3, 4, 5},
	}

	fmt.Println(arr1, arr2, arr3, arr4)
}
