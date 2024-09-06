package main

import ("fmt")

func main(){

	i := 9

	j := &i

	fmt.Println( j )

	fmt.Println( *j )

}
