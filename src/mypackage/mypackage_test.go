package mypackage_test

import (
	"testing"

	"github.com/perocha/mypackage"
)

func TestHelloWorld(t *testing.T) {
	if mypackage.HelloWorld() != "Hello world!!" {
		t.Fatal("Error message")
	}
}

func TestHello(t *testing.T) {

}

func TestRandomGreeting(t *testing.T) {
	/*
		InputValue := "Pedro"

		if helloworld.RandomGreeting() != "Hi, %v. Welcome to the party!!" {

		} else if helloworld.RandomGreeting() != "Great to see you again %v!" {

		} else if helloworld.RandomGreeting() != "Hello %v! How's it going?" {

		}
	*/
}
