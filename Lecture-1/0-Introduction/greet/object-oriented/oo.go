package main

import "fmt"

// Person struct
type Person struct {
	name string
}

// greet function with return value and interface
// this function is implemented by the Person struct

func (p Person) greet() string {
	return "Hello " + p.name + " :)"
}

func main() {
	greeter := Person{"Tohan"}
	greeting := greeter.greet()
	fmt.Printf("%s\n", greeting)
}
