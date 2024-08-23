package main

import ("fmt")

func main() {
	
	day := "Monday"

	switch day{

		case "Monday": 
			fmt.Printf("It's Monday.\n")
		case "Tuesday": 
			fmt.Printf("It's Tuesday.\n")
		case "Wednesday": 
			fmt.Printf("It's Wendneday.\n")
		case "Thursday": 
			fmt.Printf("It's Thursday.\n")
		case "Friday":
			fmt.Printf("It's Friday.\n")
		case "Saturday":
			fmt.Printf("It's Saturday.\n")
		case "Sunday":
			fmt.Printf("It's Sunday.\n")
	}
}

