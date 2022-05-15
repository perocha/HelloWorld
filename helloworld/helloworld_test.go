package helloworld_test

import (
	"testing"

	"github.com/perocha/go-hello-world/helloworld"
)

func TestHelloWorld(t *testing.T) {
	if helloworld.HelloWorld() != "Hello world!!" {
		t.Fatal("Error message")
	}
}
