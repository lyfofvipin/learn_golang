package main

import ("fmt")

func main(){

	var import = "wrong variable name because it's a keyword of the go lang"
	fmt.Println(import)

	var 1wrong_variable_name = "wrong variable name because it's starting from a numeric value"
	fmt.Println(1wrong_variable_name)

	var wrong-variable_name = "wrong variable name as it has `-` in it"
	fmt.Println(wrong-variable_name)

	var variable_name = "This is a right variable name to understand the concept of case sensitive"
	fmt.Println(variable_name)

	var variable_name = "wrong variable name as variable_name is already defined"
	fmt.Println(variable_name)
}
