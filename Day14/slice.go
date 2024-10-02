package main

import (
	"fmt"
)

func main (){

	var sli1 []int
	sli2 := []int{ 10, 20, 30, 40, 50 }

	var sli3 = []string{ "A", "a", "b", "B", `c`, `C` }
	sli4 := []rune{ 'a', '4', '-', '+' }

	sli5 := []bool { true }
	sli6 := []bool { true, false }

	fmt.Println(sli1, sli2, sli3, sli4, sli5, sli6 )

}
