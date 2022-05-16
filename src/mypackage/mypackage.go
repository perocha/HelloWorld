package mypackage

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// My first go program!
func HelloWorld() string {
	return "Hello world!!"
}

// Init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// New and more sofisticated hello function
func Hello(Name string) (string, error) {
	// If no input is given, return an error with a message
	if Name == "" {
		return "", errors.New("empty name")
	}

	// If input is received, return a value
	message := fmt.Sprintf(RandomGreeting(), Name)
	return message, nil
}

// RamdomGreeting will return a different greeting message randomly
func RandomGreeting() string {
	// A slide of message formats
	formats := []string{
		"Hi, %v. Welcome to the party!!",
		"Great to see you again %v!",
		"Hello %v! How's it going?",
	}

	// Return the available messages randomly
	return formats[rand.Intn(len(formats))]
}
