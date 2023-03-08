package main

import "fmt"

func getLengthOfCentralPark() int32 {
	var lengthInBlocks int32
	lengthInBlocks = 51
	return lengthInBlocks
	// Code after return will not be run
	fmt.Println("Whatever, I will not get printed :( !!!")
	return 0
}

// func getLengthOfCentralPark(lengthInBlocks int32) {
// 		fmt.Println("Function overloading is not supported by GOs")
// }

/*
defining parameters for a function
If the parameters are of same type, then it can also be written as:

```go

	func multiplier(x, y int32) int32 {}

```
*/
func multiplier(x int32, y int32) int32 {
	return x * y
}

func multipleValueReturn(x int, y int) (int, int) {

	// defer is used to delay the action to the end current scope, for example in this case just before return statement
	defer fmt.Println("Returning Swaped!!!")
	temp := x
	x = y
	y = temp

	// multiple vaules are returned with same return statement seperated by ","
	return x, y
}

func main() {
	var centralParkLength int32
	centralParkLength = getLengthOfCentralPark()
	fmt.Println(centralParkLength)

	// passing arguments to the function
	fmt.Println(multiplier(2, 3))

	x, y := multipleValueReturn(2, 3)
	fmt.Println(x, y)
}
