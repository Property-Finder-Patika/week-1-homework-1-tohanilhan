package main

import (
	"fmt"
	"strings"
)

func main() {
	// create greet function in Turkish
	greetPrinter(createGreetInTurkish, "Tohan")
	// create greet function in English
	greetPrinter(createGreetInEnglish, "Tohan")
	// create greet function in Turkish with uppercase
	greetPrinter(convertToUppercase, "nasılsın?")

	greetCreator := createGreetInEnglish
	greetPrinter(greetCreator, "Tohan")

	func(name string) {
		greeting := "Hello " + name + " :)"
		fmt.Printf("%s\n", greeting)
	}("Gökçe")

	closure := func(name string) {
		greeting := "Hello " + name + " :)"
		fmt.Printf("%s\n", greeting)
	}
	closure("Tuana")
	anotherGreetPrinter(closure, "Tohan")
}

// greeter function in Turkish
func createGreetInTurkish(name string) string {
	return "Merhaba " + name + " :)"
}

// greetCreator function in English
func createGreetInEnglish(name string) string {
	return "Hi " + name + " :)"
}

// convertToUppercase function converts the given string to uppercase
func convertToUppercase(arg string) string {
	return strings.ToUpper(arg)
}

// greetPrinter function prints the greeting with the given function
func greetPrinter(function func(it string) string, name string) {
	var greeting = function(name)
	fmt.Printf("%s\n", greeting)
}

// anotherGreetPrinter function prints the greeting with the given function
func anotherGreetPrinter(function func(it string), name string) {
	function(name)
}
