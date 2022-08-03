package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Chris")
	fmt.Println(message)
}
