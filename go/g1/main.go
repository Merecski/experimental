package main

import (
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func main() {
	msg := Hello("John")
	fmt.Printf("%s\n", msg)
	x := 5 - 10
	fmt.Println("Variable x contains ", x, " as it's value.")
}
