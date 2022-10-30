package main

// Use import to use functions from other packages
/***
  You can also use import as:
  ```go
  import "fmt"
  import "rsc.io/quote"
  ```
*/
import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	// Get a greeting message and print it.
	// fmt.Println(message)
	fmt.Println(quote.Go())
}
