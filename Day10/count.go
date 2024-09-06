package main

import ("fmt")

func main (){

	for i := 1 ; i <= 10; i = i +1 {
		fmt.Printf( "%d \n" , i )
	}
	// fmt.Println(i) // this is a wrong statement as i is not accessible after the loop
}