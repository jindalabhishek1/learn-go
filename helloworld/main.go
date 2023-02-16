package main

/*
* package is used to group .go files with similar functionality
* This tells go compiler, to which package this .go file belongs
* In context to this file, it means, main.go belongs to main package
 */

/*
* import statement allows to use code from other packages
* we can import multiple packages
	* option1:
		* import "package1"
		* import "package2"
	* option2:
		* import (
			p1 "package1"
			"package2"
		)
		* p1 is the alias for the package
*/
import "fmt"

/*
* func keyword: used to declare Go function
* code should be indented inside the curly braces
* main function in main package is starting point in go
 */
func main() {
	fmt.Println("Hello World")
}
