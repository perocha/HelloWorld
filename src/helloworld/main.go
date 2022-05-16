package main

import (
	"fmt"
	"strconv"

	"github.com/perocha/calculator"

	"github.com/perocha/mypackage"

	"log"

	"rsc.io/quote"
)

func main() {
	// Print something
	firstName, lastName := "John", "Doe"
	age := 32
	fmt.Println(firstName, lastName, age)

	// Print something more
	fullName := firstName + " " + lastName + "\t Age: " + strconv.Itoa(age)
	fmt.Println(fullName)

	// Perform some calculations
	resultSum, resultMul := SomeCalculations("150", "123")
	fmt.Println("Sum ", resultSum)
	fmt.Println("Mul ", resultMul)

	// Example using pointers
	firstName = "John"
	fmt.Println(firstName)
	updateName(&firstName)
	fmt.Println(firstName)

	//
	resultCalc := calculator.CalculatorSum(120, 123)
	fmt.Println("Result: ", resultCalc)

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Basic hello world message
	fmt.Println(mypackage.HelloWorld())
	fmt.Println(quote.Go())

	// Request a greeting message
	message, err := mypackage.Hello("Pedro")
	// If error is received, print the error in the console and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console
	fmt.Println(message)
}

func SomeCalculations(InputNumber1 string, InputNumber2 string) (sum int, mul int) {
	intInputNumber1, _ := strconv.Atoi(InputNumber1)
	intInputNumber2, _ := strconv.Atoi(InputNumber2)
	sum = intInputNumber1 + intInputNumber2
	mul = intInputNumber1 * intInputNumber2
	return
}

func updateName(InputName *string) {
	*InputName = "David"
}
