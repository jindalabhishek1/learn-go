package greetings

import "fmt"

// Hello returns s greeting for the named person.
// Syntax: func <name of the function> (<parameter name> <parameter type>) <return type> {
func Hello(name string) string {
    // Returns a greeting that embeds the name in a message.
    // Syntax: <variable name> := <something>
    // := is shortcut for declaring and initializing a variable
    // Same as: var message string
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
