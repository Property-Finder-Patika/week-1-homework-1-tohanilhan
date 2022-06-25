package main

import (
	"fmt"
	"os"
)

func main() {
	// prints the string directly
	greet("Tohan") // Hello Tohan :)

	// prints the string returned by the function
	greetings := greetWithReturn("Tohan")
	fmt.Println(greetings) // Hello Tohan :)

	// prints the string created by the function with argument
	// you can run the program with the argument "Tohan" to see the result
	// example: go run greet_with_function.go Tohan
	name := os.Args[1]
	greeting := createGreet(name)
	fmt.Printf("%s\n", greeting)

	// we can also use the scanln function to get the input from the user instead of using os.Args
	var anotherName string
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)
	greetingWithAnotherName := createGreet(anotherName)
	fmt.Printf("%s\n", greetingWithAnotherName)

}

// greet function without return value
func greet(name string) {
	fmt.Printf("Hello %s :)\n", name)
}

// greeting function with return value
func greetWithReturn(name string) string {
	greeting := "Hello " + name + " :)"
	return greeting
}

// createGreet function with return value created with argument
func createGreet(name string) string {
	return "Hello " + name + " :)"
}
