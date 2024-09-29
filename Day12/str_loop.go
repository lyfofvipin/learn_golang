package main

import "fmt"

func main() {
	
	name := "Vipin Kumar"

	for str := range name {

		fmt.Printf( "Index %d has value %c.\n", str, name[str]  )

	}

}
