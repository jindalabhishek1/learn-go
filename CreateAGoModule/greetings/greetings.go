package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
// Syntax: func <name of the function> (<parameter name> <parameter type>) <return type> {
// Hint: Functions starting with a capital letter can be called by a function not in the same package.
//
//	This is known as exported name.
func Hello(name string) (string, error) {
	// Returns a greeting that embeds the name in a message.

	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name in a greeting message.

	// Syntax: <variable name> := <something>
	// := is shortcut for declaring and initializing a variable
	// Same as: var message string
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat())
	return message, nil
}

// init sets initial values for variables used in the function.
// Go executes the init functions automatically at program startup, after global variables have been initialized.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
// Hint: funtion starting with lowercase letter is not exported and are accessible to the code in its own package
func randomFormat() string {
	// A slice of message formates.
	// Slice is a kind of dynamic array in Go.
	// The ',' at the end of last element in slice is needed if the closing bracket is in next line (after a new line character)
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	// Syntax for a map: make(map[key-type]value-type)
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	// range returns two values: the index of the current item in the loop and a copy of the item's value.
	// _ (the Go blank identifier, an underscore) is used to ignore something.
	// here we don't need the index, hence ignoring it.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}

	return messages, nil
}
