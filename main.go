package main

import (
	"fmt"

	"log"

	"github.com/perocha/go-hello-world/helloworld"

	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Basic hello world message
	fmt.Println(helloworld.HelloWorld())
	fmt.Println(quote.Go())

	// Request a greeting message
	message, err := helloworld.Hello("Pedro")
	// If error is received, print the error in the console and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console
	fmt.Println(message)
}
