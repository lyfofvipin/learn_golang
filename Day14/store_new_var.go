package main

import "fmt"

func main(){

	sli1 := []int{ 1, 2, 3, 4 }
	sli2 := []int{ 10, 20, 30, 40 }

	fmt.Println( sli1, sli2 )
	sli1 = append( sli1, sli2... )
	fmt.Println( sli1, sli2 )

}
