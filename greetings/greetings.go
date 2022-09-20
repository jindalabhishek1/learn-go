package greetings

import (
    "errors"
    "fmt"
)

// Hello returns s greeting for the named person.
// Syntax: func <name of the function> (<parameter name> <parameter type>) <return type> {
func Hello(name string) (string, error) {
    // Returns a greeting that embeds the name in a message.

    // If no name was given, return an error with a message.
    if name == "" {
        retun "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name in a greeting message.

    // Syntax: <variable name> := <something>
    // := is shortcut for declaring and initializing a variable
    // Same as: var message string
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
