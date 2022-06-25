package main

import "fmt"

func main() {
	greet("Tohan")
}

func greet(name string) {
	fmt.Printf("Hello %s :)\n", name)
}
