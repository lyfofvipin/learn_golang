package main

import ("fmt")

func main() {

	value1 := 5
	value2 := 55
	operator := ""

	fmt.Println("Enter a value from these operators +, -, *,/ : ")
	fmt.Scan(&operator)

	switch operator {

	case "+":
		fmt.Println( value1 + value2 )
	case "-":
		fmt.Println( value1 - value2 )
	case "*":
		fmt.Println( value1 * value2 )
	case "/":
		fmt.Println( value1 / value2 )
	default:
		fmt.Println("Seems like you are entering some unsupported operator.")
	}

}