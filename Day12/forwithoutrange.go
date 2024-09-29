package main

import "fmt"

func main() {
	
	name := "Vipin Kumar"

	length := len(name)

	for iterator := 0 ; iterator < length ; iterator += 1 {

		fmt.Printf( "Index Value %d has character %c.\n", iterator, name[iterator] )

	}

}
