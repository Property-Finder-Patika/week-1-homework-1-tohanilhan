package main

import "fmt"

type Person struct {
	name                             string
}

// when you save the file, it will format the code
// or you can run the file with the command: go fmt greet.go to format the code

func (p Person) greet()                             string {
	return         "Hello " +          p.name + " :)"
}

func main() {
	var greeter              Person =            Person{"Tohan"}

	greeting := greeter.greet()
	fmt.Printf("%s\n", greeting)
}
