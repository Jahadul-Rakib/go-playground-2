package main

import "fmt"

func main() {

	fmt.Print("From Pacage....")

	printSome("Print From Function")

	result := getResult(2, 3)
	println("Result is : ", result)

	addResult, subTractResult := getResultMultipleReturnType(20, 12)
	println("add result: ", addResult, " sub result: ", subTractResult)

	addResult1, subTractResult1 := getResultMultipleReturnTypeAnotherType(30, 23)
	println("add result: ", addResult1, " sub result: ", subTractResult1)

	anonymousFunction := func(a string) {
		println("From anonymous function: ", a)
	}
	anonymousFunction("call")

}
func printSome(data string) {
	println(data)
}

func getResult(a int, b int) int {
	return a + b
}
func getResultMultipleReturnType(a int, b int) (int, int) {
	add := a + b
	subtract := a - b
	return add, subtract
}

func getResultMultipleReturnTypeAnotherType(a int, b int) (add, subtract int) {
	add = a + b
	subtract = a - b
	return
}
