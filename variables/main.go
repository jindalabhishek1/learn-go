package main

import "fmt"

func main() {

	/*
		* Literals: values which are literally what they say.
		  For e.g. numbers or text wrapped in quotes
	*/
	fmt.Println(2 * 10.5)
	fmt.Println("Hello" + "World")
	// fmt.Println("Hello" * "World") // doesn't work

	// const keyword is used to create constant named values
	const constVariable = 10

	/*
		* Data type: Designation by a programming language about what information is being stored
		* Numbers:
			* Types:
				* Int: 1, -2343, 0
				* Float: 1.343, -345.324
				* Complex (containing imaginary part): 2i, -34+0.3i
			* 15 different numeric types for above mentioned 3 types (how many bits used to store)
				* 11 for int, 2 for float, 2 for complex
				* link: https://go.dev/ref/spec#Numeric_types
			* Integers:
				* Signed: can be nagative
				* UnSigned: can only be positive, including 0
		* Boolean:
			* True
			* False
	*/

	var someVarWithTypeSpecified string = "String"
	fmt.Println(someVarWithTypeSpecified)

	someVarWithNoTypeSpecified := 1232

	fmt.Println(someVarWithNoTypeSpecified)

	someVarWithNoTypeSpecifiedExample2 := "abf"

	fmt.Println(someVarWithNoTypeSpecifiedExample2)

	a, c := "abc", 1

	fmt.Println(a, c)

}
