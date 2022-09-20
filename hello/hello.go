package main

import (
     "fmt"

    "github.com/jindalabhishek1/learn-go/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello ("Gladys")
    fmt.Println (message)
}
