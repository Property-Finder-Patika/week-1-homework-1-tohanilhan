package main

import "fmt"

func main() {
	// prints the string directly
	greet("Tohan") // Hello Tohan :)

	// prints the string returned by the function
	greetings := greetWithReturn("Tohan")
	fmt.Println(greetings) // Hello Tohan :)
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
